# Report Scan Edge Case, Query, dan Fitur Komentar

Tanggal scan: 2026-05-16  
Scope: repository SQL, migration, handler/service terkait `permintaan`, `distribusi`, `pelaksana`, `laporan`, `verifikasi`, `all_activity`, dan `superadmin_dashboard`.

## Ringkasan Eksekutif

Kode sudah memakai query parameter untuk input user langsung, jadi risiko SQL injection utama relatif terkendali. Masalah yang lebih menonjol ada di konsistensi data, query list yang berpotensi berat saat data membesar, dan fitur komentar yang masih setengah matang: belum ada validasi isi komentar, belum ada update/delete/read state yang lengkap, response JSON belum konsisten, serta `is_read` komentar belum dipakai.

Prioritas tertinggi:

1. Tambahkan transaksi pada operasi multi-step `CreateDistribusi`, `UpdateDistribusi`, `CreateLaporan`, dan `UpdateLaporan`.
2. Tambahkan pagination dan index untuk endpoint list/dashboard/activity.
3. Matangkan fitur komentar: validasi required/min/max, author id di response, mark-read per penerima, update/delete, dan naming field konsisten.
4. Perbaiki bug query `UpdateLampiran` laporan yang mengarah ke tabel `laporan_progress`, padahal tabel yang ada adalah `laporan_kinerja`.

## Temuan Edge Case dan Bug

### 1. Operasi multi-write belum atomic

Lokasi:

- `internal/distribusi/service.go:18-26`
- `internal/distribusi/service.go:62-70`
- `internal/laporan/service.go:37-53`
- `internal/laporan/service.go:116-142`

Contoh risiko:

- Distribusi berhasil dibuat, tetapi insert pelaksana gagal. Hasilnya ada row `distribusi` tanpa pelaksana.
- Update distribusi berhasil, lalu `ReplacePelaksana` gagal setelah delete pelaksana lama. Hasilnya assignment pelaksana bisa hilang.
- Laporan berhasil dibuat, tetapi history gagal dibuat. Audit trail bolong.
- Laporan berhasil di-update, status verifikasi atau history gagal. State antar tabel tidak sinkron.

Rekomendasi:

- Gunakan `db.BeginTx`, lalu pass `*sql.Tx` ke repository.
- Commit hanya setelah semua step berhasil.
- Rollback otomatis jika salah satu step gagal.

### 2. `UpdateLampiran` laporan memakai nama tabel yang salah

Lokasi:

- `internal/laporan/repository.go:597-599`
- Migration tabel laporan: `migrations/000008_create_laporan_kinerja_table.up.sql:1`
- Kolom lampiran ditambahkan ke `laporan_kinerja`: `migrations/000017_add_lampiran_laporan_kinerja_table.up.sql:1`

Query saat ini:

```sql
UPDATE laporan_progress SET lampiran = $1, updated_at = NOW() WHERE id = $2
```

Tabel `laporan_progress` tidak ada di migration. Endpoint upload lampiran laporan berpotensi selalu gagal di runtime.

Rekomendasi:

- Ganti menjadi `UPDATE laporan_kinerja ...`.
- Tambahkan test/integration test untuk `PATCH /laporan/:id/lampiran`.

### 3. `GetById` mengabaikan error pertama pada batch fetch

Lokasi:

- `internal/distribusi/repository.go:128-130`
- `internal/laporan/repository.go:155-157`

Pola saat ini:

```go
pelaksanaMap, err := getPelaksanaByDistribusiIDs(...)
komentarMap, err := getKomentarByDistribusiIDs(...)
if err != nil { ... }
```

Jika query pertama gagal tapi query kedua berhasil, error pertama tertimpa dan tidak pernah dikembalikan. Pola yang sama terjadi di laporan untuk `getVerifikasiByLaporanIDs`.

Rekomendasi:

- Cek error langsung setelah setiap call.

### 4. Beberapa query list tidak memanggil `rows.Err()`

Lokasi contoh:

- `internal/distribusi/repository.go:33-85`
- `internal/distribusi/repository.go:243-277`
- `internal/pelaksana/repository.go:18-29`
- `internal/superadmin_dashboard/repository.go:141-159`, `184-192`, `222-238`
- `internal/laporan/repository.go:292-312`

Jika error terjadi saat iterasi row, fungsi bisa mengembalikan data parsial tanpa error.

Rekomendasi:

- Setelah loop `for rows.Next()`, selalu panggil `if err := rows.Err(); err != nil { return ..., err }`.

### 5. Update dengan `QueryRow(... RETURNING ...)` sudah bagus, tapi pesan error bisa misleading

Lokasi:

- `internal/distribusi/repository.go:405-424`
- `internal/laporan/repository.go:571-594`
- `internal/permintaan/repository.go:185-218`

Jika `id` tidak ditemukan, `sql.ErrNoRows` akan dikembalikan. Handler sudah mengubah menjadi 404. Namun untuk FK error di distribusi/laporan, pesan error hanya menyebut `permintaan_id`, padahal bisa juga `admin_id`, `programmer_id`, atau `verifikator_id` tergantung query.

Rekomendasi:

- Mapping error FK sebaiknya spesifik berdasarkan constraint name `pqErr.Constraint`, bukan hanya code `23503`.

## Query yang Kemungkinan Bermasalah

### 1. Endpoint list dan dashboard belum punya pagination

Lokasi:

- `internal/distribusi/repository.go:10-27`
- `internal/laporan/repository.go:10-26`
- `internal/permintaan/repository.go:46-61`
- `internal/all_activity/repository.go:8-116`
- `internal/superadmin_dashboard/repository.go:11-115`

Dampak:

- Semua data di-load sekaligus.
- Dashboard superadmin mengambil semua permintaan, lalu batch fetch distribusi, pelaksana, dan laporan.
- `all_activity` melakukan `UNION ALL` semua aktivitas lalu `ORDER BY created_at DESC` tanpa limit.

Rekomendasi:

- Tambahkan `limit`, `offset` atau cursor pagination.
- Untuk dashboard, pertimbangkan endpoint ringkasan terpisah dari detail list.
- Default limit, misalnya 20 atau 50, dengan max limit.

### 2. Dynamic `IN ($1,$2,...)` bisa membesar dan mengubah query plan

Lokasi:

- `internal/distribusi/repository.go:156-170`, `192-207`
- `internal/laporan/repository.go:259-275`, `321-336`
- `internal/pelaksana/repository.go:150-165`
- `internal/superadmin_dashboard/repository.go:117-132`, `162-178`, `195-216`

Risiko:

- Saat jumlah ID besar, SQL string dan jumlah parameter ikut besar.
- PostgreSQL bisa menghasilkan plan berbeda-beda karena query text berubah.
- Ada batas parameter PostgreSQL dan overhead parsing meningkat.

Rekomendasi:

- Gunakan `WHERE id = ANY($1)` dengan `pq.Array(ids)` atau `uuid[]`.
- Kombinasikan dengan pagination agar jumlah ID batch tetap kecil.

### 3. `DISTINCT ON` untuk status verifikasi terbaru belum punya tie-breaker kuat

Lokasi:

- `internal/distribusi/repository.go:12-27`
- `internal/distribusi/repository.go:92-109`
- `internal/laporan/repository.go:266-273`
- `internal/verifikasi/repository.go` juga memakai pola sejenis.

Masalah:

- `ORDER BY ... updated_at DESC` saja tidak deterministik jika dua row punya `updated_at` sama.
- `DISTINCT ON (d.id)` pada distribusi mengambil satu verifikasi dari join laporan/verifikasi. Jika satu permintaan punya beberapa laporan dan beberapa verifikasi, hasil yang tampil hanya satu status terakhir global, bukan struktur lengkap.

Rekomendasi:

- Tambahkan tie-breaker `ORDER BY ..., updated_at DESC, id DESC`.
- Jika yang dibutuhkan status per laporan, jangan flatten ke satu verifikasi di distribusi. Kembalikan array atau ambil latest per laporan dengan subquery/window function.

### 4. Banyak `LEFT JOIN` ke tabel yang secara domain wajib ada

Lokasi contoh:

- `internal/distribusi/repository.go:20-25`
- `internal/laporan/repository.go:20-24`
- `internal/permintaan/repository.go:56-59`

Karena migration memakai FK `NOT NULL`, relasi utama seperti `permintaan`, `users`, `master_pemda`, `master_aplikasi` seharusnya ada. `LEFT JOIN` membuat data korup terlihat sebagai field kosong/zero value, dan bisa menambah ambiguitas ketika scan ke string non-nullable.

Rekomendasi:

- Gunakan `INNER JOIN` untuk relasi wajib.
- Pakai `LEFT JOIN` hanya untuk relasi opsional seperti verifikasi/komentar bila memang bisa kosong.

### 5. `all_activity` berpotensi mahal

Lokasi:

- `internal/all_activity/repository.go:8-116`

Query melakukan `UNION ALL` lima sumber aktivitas, join ke beberapa tabel, lalu sort seluruh hasil. Tanpa pagination dan index pendukung, endpoint ini akan menjadi salah satu bottleneck lebih awal.

Rekomendasi:

- Tambahkan `LIMIT/OFFSET` atau cursor berbasis `created_at`.
- Tambahkan index `created_at` di tabel sumber.
- Pertimbangkan materialized view atau table activity log bila fitur ini sering diakses.

## Query yang Bisa Dioptimalkan

### Index yang disarankan

Saat ini migration hampir tidak membuat index eksplisit selain primary key dan unique. FK di PostgreSQL tidak otomatis membuat index di child table.

Tambahkan index untuk kolom yang sering dipakai di `WHERE`, `JOIN`, dan `ORDER BY`:

```sql
CREATE INDEX idx_permintaan_is_archived_created_at ON permintaan (is_archived, created_at DESC);
CREATE INDEX idx_distribusi_permintaan_id_created_at ON distribusi (permintaan_id, created_at DESC);
CREATE INDEX idx_distribusi_admin_id ON distribusi (admin_id);
CREATE INDEX idx_distribusi_pelaksana_programmer_id_created_at ON distribusi_pelaksana (programmer_id, created_at DESC);
CREATE INDEX idx_laporan_kinerja_permintaan_id_created_at ON laporan_kinerja (permintaan_id, created_at DESC);
CREATE INDEX idx_laporan_kinerja_programmer_id_created_at ON laporan_kinerja (programmer_id, created_at DESC);
CREATE INDEX idx_verifikasi_laporan_id_updated_at ON verifikasi (laporan_id, updated_at DESC, id DESC);
CREATE INDEX idx_komentar_distribusi_distribusi_id_created_at ON komentar_distribusi (distribusi_id, created_at ASC);
CREATE INDEX idx_komentar_laporan_laporan_id_created_at ON komentar_laporan (laporan_id, created_at ASC);
CREATE INDEX idx_progress_history_laporan_id_created_at ON progress_history (laporan_id, created_at DESC);
```

### Batch fetch bisa dibuat lebih sederhana

Pola batch fetch sudah lebih baik daripada N+1 query, tetapi masih membuat placeholder manual. Untuk readability dan plan stability:

```sql
WHERE kd.distribusi_id = ANY($1)
```

dengan `pq.Array(ids)`.

### Dashboard bisa dipecah

Lokasi:

- `internal/superadmin_dashboard/repository.go:11-115`

Endpoint dashboard saat ini mengembalikan semua detail permintaan, distribusi, pelaksana, dan laporan. Jika UI hanya perlu angka total di awal, query detail ini terlalu berat.

Rekomendasi:

- Endpoint `/superadmin-dashboard/summary`: count per status/tabel.
- Endpoint `/superadmin-dashboard/items`: paginated list.
- Detail per permintaan diambil lazy saat user membuka item.

## Fitur Komentar yang Belum Mature

### 1. Field dan naming belum konsisten

Lokasi:

- Request distribusi: `internal/distribusi/dto.go:20-22` memakai `json:"komentars"`.
- Response list distribusi: `internal/distribusi/dto.go:107-112` field Go `Komentars`, JSON `komentar`.
- Response create komentar distribusi: `internal/distribusi/dto.go:135-140` JSON `komentars`.
- Tabel distribusi: `migrations/000019_create_komentar_dsitribusi_table.up.sql:5` kolom `komentars`.
- Laporan memakai `komentar` di request/table/response.

Dampak:

- Frontend harus mengingat variasi `komentar` vs `komentars`.
- Nama `komentars` tidak natural dan rawan salah mapping.

Rekomendasi:

- Standarkan ke `komentar` untuk isi satu komentar.
- Gunakan `komentar_list` atau `komentars` hanya untuk array response jika memang sudah menjadi kontrak API, tapi sebaiknya konsisten di semua modul.
- Tambahkan migration rename kolom `komentars` ke `komentar` jika API belum stabil.

### 2. Komentar tidak wajib dan tidak punya batas panjang

Lokasi:

- `internal/distribusi/dto.go:20-22`
- `internal/laporan/dto.go:30-32`

Saat ini request komentar tidak punya `validate:"required,min=1,max=..."`. Komentar kosong bisa masuk DB karena kolom `TEXT` nullable.

Rekomendasi:

- DTO: `validate:"required,min=1,max=2000"`.
- Migration: `komentar TEXT NOT NULL CHECK (length(trim(komentar)) > 0)`.
- Normalisasi trim whitespace di service.

### 3. `is_read` komentar disimpan tapi belum dipakai

Lokasi:

- `migrations/000019_create_komentar_dsitribusi_table.up.sql:6`
- `migrations/000020_create_komentar_laporan_table.up.sql:6`
- Insert mengembalikan `is_read`: `internal/distribusi/repository.go:383-400`, `internal/laporan/repository.go:549-566`

Belum terlihat endpoint untuk mark-read komentar, unread count, atau read state per penerima. Karena `is_read` berada di row komentar, bukan per user, satu boolean juga tidak cukup jika komentar bisa dibaca oleh banyak role/user.

Rekomendasi:

- Jika komentar hanya punya satu penerima, simpan `recipient_user_id` dan mark-read spesifik.
- Jika komentar dibaca banyak user, buat tabel `komentar_reads(comment_id, user_id, read_at)`.
- Tambahkan endpoint `PATCH /.../komentar/:id/read` dan unread count.

### 4. Response komentar tidak menyertakan `user_id`/author id

Lokasi:

- `internal/distribusi/repository.go:199-216`
- `internal/laporan/repository.go:328-345`
- DTO hanya punya `full_name`: `internal/distribusi/dto.go:107-112`, `internal/laporan/dto.go:84-89`

Dampak:

- Frontend sulit membedakan komentar milik user login.
- Tidak bisa membuat action edit/delete hanya untuk pemilik komentar.
- Nama user tidak stabil sebagai identifier.

Rekomendasi:

- Include `user_id`, `username`, dan optional `profile_picture` di response komentar.

### 5. Belum ada update/delete komentar

Route hanya create:

- `routes/routes.go:96`
- `routes/routes.go:121`

Belum ada:

- `GET /komentar/:id`
- `PATCH /komentar/:id`
- `DELETE /komentar/:id`
- authorization bahwa hanya author atau role tertentu yang boleh edit/delete.

Rekomendasi:

- Tambahkan endpoint CRUD minimal untuk komentar.
- Simpan `updated_at` saat komentar diedit.
- Pertimbangkan soft delete jika komentar dipakai sebagai audit trail.

### 6. Authorization komentar masih terlalu longgar

Lokasi:

- Group distribusi mengizinkan `super_admin`, `admin`, `programmer`: `routes/routes.go:87-96`
- Group laporan mengizinkan `programmer`, `verifikator`: `routes/routes.go:109-121`

Handler hanya validasi JWT dan UUID, lalu insert komentar. Belum ada validasi bahwa user memang terkait dengan distribusi/laporan tersebut.

Contoh risiko:

- Programmer A bisa komentar di distribusi milik Programmer B jika tahu `distribusi_id`.
- Programmer bisa komentar ke laporan yang bukan miliknya jika tahu `laporan_id`.

Rekomendasi:

- Validasi membership sebelum insert:
  - Distribusi: user adalah admin pembuat, pelaksana terkait, atau super_admin.
  - Laporan: user adalah programmer pemilik laporan atau verifikator terkait.

### 7. Komentar distribusi lama dan thread komentar baru bercampur konsep

Lokasi:

- Tabel `distribusi` punya kolom `komentar`: `migrations/000006_create_distribusi_table.up.sql:5`
- Tabel `komentar_distribusi` punya komentar thread: `migrations/000019_create_komentar_dsitribusi_table.up.sql:1-9`
- Response detail distribusi memakai `d.komentar`: `internal/distribusi/repository.go:224-253`
- Response full distribusi memakai `komentar_distribusi`: `internal/distribusi/repository.go:75-82`

Dampak:

- Ada dua sumber komentar distribusi dengan makna berbeda.
- UI/API bisa bingung: komentar awal distribusi vs thread komentar.

Rekomendasi:

- Rename konsep:
  - `distribusi.catatan_admin` untuk catatan awal assignment.
  - `komentar_distribusi.komentar` untuk thread.
- Dokumentasikan keduanya di Swagger.

## Catatan Migration

### Nama migration typo

Lokasi:

- `migrations/000019_create_komentar_dsitribusi_table.up.sql`
- `migrations/000019_create_komentar_dsitribusi_table.down.sql`

Typo `dsitribusi` tidak merusak runtime jika sudah konsisten, tetapi mengganggu maintenance dan pencarian.

### Tidak ada trigger `updated_at`

Banyak tabel punya `updated_at DEFAULT NOW()`, tetapi update timestamp bergantung pada setiap query update manual. Untuk komentar, belum ada update endpoint, tapi jika nanti ada, mudah lupa set `updated_at`.

Rekomendasi:

- Tambahkan trigger umum `set_updated_at()` untuk tabel yang punya `updated_at`.

## Rekomendasi Roadmap Perbaikan

### Quick wins

1. Perbaiki `UPDATE laporan_progress` menjadi `UPDATE laporan_kinerja`.
2. Tambahkan `rows.Err()` di semua repository list/helper.
3. Cek error tiap batch call sebelum lanjut ke batch berikutnya.
4. Tambahkan validasi required/min/max untuk request komentar.
5. Tambahkan index untuk FK dan query list utama.

### Menengah

1. Refactor dynamic `IN` menjadi `ANY($1)` dengan array parameter.
2. Tambahkan pagination ke endpoint list, dashboard, dan all-activity.
3. Tambahkan transaksi untuk create/update distribusi dan laporan.
4. Tambahkan tie-breaker `id DESC` di query `DISTINCT ON`.

### Fitur komentar mature

1. Standarkan naming `komentar`.
2. Tambahkan author metadata di response.
3. Tambahkan authorization berdasarkan relasi user ke distribusi/laporan.
4. Tambahkan update/delete komentar.
5. Desain ulang read state komentar: per penerima atau tabel read receipt.
6. Tambahkan test untuk komentar kosong, unauthorized comment, komentar milik user lain, dan unread/read flow.

