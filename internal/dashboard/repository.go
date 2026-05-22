package dashboard

import (
	"aplikasi-internal/config"
	"fmt"
	"strings"
	"time"
)

// ── Helpers ───────────────────────────────────────────────────────────────────

func makePlaceholders(n int) string {
	ph := make([]string, n)
	for i := range ph {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	return strings.Join(ph, ",")
}

func toArgs(ids []string) []interface{} {
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	return args
}

// ── SuperAdmin ────────────────────────────────────────────────────────────────

func fetchAllPermintaan() ([]PermintaanItem, error) {
	rows, err := config.DB.Query(`
		SELECT
			p.id,
			mp.id, mp.name, mp.logo,
			ma.id, ma.name, ma.logo,
			p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran, p.status,
			u.id, u.username, u.full_name, u.profile_picture,
			p.created_at, p.updated_at
		FROM permintaan p
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON p.created_by = u.id
		ORDER BY p.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PermintaanItem
	for rows.Next() {
		var item PermintaanItem
		if err := rows.Scan(
			&item.ID,
			&item.Pemda.ID, &item.Pemda.Name, &item.Pemda.Logo,
			&item.Aplikasi.ID, &item.Aplikasi.Name, &item.Aplikasi.Logo,
			&item.Menu, &item.KondisiAwal, &item.KondisiDiharapkan,
			&item.TanggalPesanan, &item.TanggalDeadline, &item.Lampiran, &item.Status,
			&item.Pembuat.ID, &item.Pembuat.Username, &item.Pembuat.FullName, &item.Pembuat.ProfilePicture,
			&item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		item.Distribusi = []DistribusiItem{}
		item.Laporan = []LaporanItem{}
		items = append(items, item)
	}
	return items, rows.Err()
}

func fetchDistribusiByPermintaanIDs(permintaanIDs []string) (map[string][]DistribusiItem, []string, error) {
	query := fmt.Sprintf(`
		SELECT
			d.id, d.permintaan_id,
			u.id, u.username, u.full_name, u.profile_picture,
			d.komentar, d.created_at, d.updated_at
		FROM distribusi d
		LEFT JOIN users u ON d.admin_id = u.id
		WHERE d.permintaan_id IN (%s)
		ORDER BY d.created_at ASC
	`, makePlaceholders(len(permintaanIDs)))

	rows, err := config.DB.Query(query, toArgs(permintaanIDs)...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	result := make(map[string][]DistribusiItem)
	var distribusiIDs []string

	for rows.Next() {
		var permID string
		var d DistribusiItem
		var createdAt, updatedAt time.Time
		if err := rows.Scan(
			&d.ID, &permID,
			&d.Admin.ID, &d.Admin.Username, &d.Admin.FullName, &d.Admin.ProfilePicture,
			&d.Komentar, &createdAt, &updatedAt,
		); err != nil {
			return nil, nil, err
		}
		d.CreatedAt = createdAt
		d.UpdatedAt = updatedAt
		d.Pelaksana = []UserInfo{}
		result[permID] = append(result[permID], d)
		distribusiIDs = append(distribusiIDs, d.ID)
	}
	return result, distribusiIDs, rows.Err()
}

func fetchPelaksanaByDistribusiIDs(distribusiIDs []string) (map[string][]UserInfo, error) {
	result := make(map[string][]UserInfo)
	if len(distribusiIDs) == 0 {
		return result, nil
	}

	query := fmt.Sprintf(`
		SELECT dp.distribusi_id, u.id, u.username, u.full_name, u.profile_picture
		FROM distribusi_pelaksana dp
		LEFT JOIN users u ON dp.programmer_id = u.id
		WHERE dp.distribusi_id IN (%s)
	`, makePlaceholders(len(distribusiIDs)))

	rows, err := config.DB.Query(query, toArgs(distribusiIDs)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var distribusiID string
		var u UserInfo
		if err := rows.Scan(&distribusiID, &u.ID, &u.Username, &u.FullName, &u.ProfilePicture); err != nil {
			return nil, err
		}
		result[distribusiID] = append(result[distribusiID], u)
	}
	return result, rows.Err()
}

func fetchLaporanByPermintaanIDs(permintaanIDs []string) (map[string][]LaporanItem, int, error) {
	result := make(map[string][]LaporanItem)
	if len(permintaanIDs) == 0 {
		return result, 0, nil
	}

	query := fmt.Sprintf(`
		SELECT
			l.id, l.permintaan_id,
			u.id, u.username, u.full_name, u.profile_picture,
			l.laporan_progress, l.status,
			l.created_at, l.updated_at
		FROM laporan_kinerja l
		LEFT JOIN users u ON l.programmer_id = u.id
		WHERE l.permintaan_id IN (%s)
		ORDER BY l.created_at ASC
	`, makePlaceholders(len(permintaanIDs)))

	rows, err := config.DB.Query(query, toArgs(permintaanIDs)...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	total := 0
	for rows.Next() {
		var permID string
		var l LaporanItem
		if err := rows.Scan(
			&l.ID, &permID,
			&l.Programmer.ID, &l.Programmer.Username, &l.Programmer.FullName, &l.Programmer.ProfilePicture,
			&l.LaporanProgress, &l.Status,
			&l.CreatedAt, &l.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		result[permID] = append(result[permID], l)
		total++
	}
	return result, total, rows.Err()
}

// ── Admin ─────────────────────────────────────────────────────────────────────

func fetchPermintaanByIDs(ids []string) ([]PermintaanItem, error) {
	if len(ids) == 0 {
		return []PermintaanItem{}, nil
	}
	query := fmt.Sprintf(`
		SELECT
			p.id,
			mp.id, mp.name, mp.logo,
			ma.id, ma.name, ma.logo,
			p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran, p.status,
			u.id, u.username, u.full_name, u.profile_picture,
			p.created_at, p.updated_at
		FROM permintaan p
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON p.created_by = u.id
		WHERE p.id IN (%s)
		ORDER BY p.created_at DESC
	`, makePlaceholders(len(ids)))

	rows, err := config.DB.Query(query, toArgs(ids)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PermintaanItem
	for rows.Next() {
		var item PermintaanItem
		if err := rows.Scan(
			&item.ID,
			&item.Pemda.ID, &item.Pemda.Name, &item.Pemda.Logo,
			&item.Aplikasi.ID, &item.Aplikasi.Name, &item.Aplikasi.Logo,
			&item.Menu, &item.KondisiAwal, &item.KondisiDiharapkan,
			&item.TanggalPesanan, &item.TanggalDeadline, &item.Lampiran, &item.Status,
			&item.Pembuat.ID, &item.Pembuat.Username, &item.Pembuat.FullName, &item.Pembuat.ProfilePicture,
			&item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		item.Distribusi = []DistribusiItem{}
		item.Laporan = []LaporanItem{}
		items = append(items, item)
	}
	return items, rows.Err()
}

func fetchDistribusiByAdminID(adminID string) (map[string][]DistribusiItem, []string, error) {
	rows, err := config.DB.Query(`
		SELECT
			d.id, d.permintaan_id,
			u.id, u.username, u.full_name, u.profile_picture,
			d.komentar, d.created_at, d.updated_at
		FROM distribusi d
		LEFT JOIN users u ON d.admin_id = u.id
		WHERE d.admin_id = $1
		ORDER BY d.created_at ASC
	`, adminID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	result := make(map[string][]DistribusiItem)
	var distribusiIDs []string

	for rows.Next() {
		var permID string
		var d DistribusiItem
		var createdAt, updatedAt time.Time
		if err := rows.Scan(
			&d.ID, &permID,
			&d.Admin.ID, &d.Admin.Username, &d.Admin.FullName, &d.Admin.ProfilePicture,
			&d.Komentar, &createdAt, &updatedAt,
		); err != nil {
			return nil, nil, err
		}
		d.CreatedAt = createdAt
		d.UpdatedAt = updatedAt
		d.Pelaksana = []UserInfo{}
		result[permID] = append(result[permID], d)
		distribusiIDs = append(distribusiIDs, d.ID)
	}
	return result, distribusiIDs, rows.Err()
}

// ── Programmer ────────────────────────────────────────────────────────────────

func fetchPenugasanByProgrammerID(programmerID string) ([]PenugasanItem, error) {
	rows, err := config.DB.Query(`
		SELECT
			p.id, p.distribusi_pelaksana_id,
			p.judul, p.deskripsi, p.deadline, p.prioritas,
			p.estimasi_hari, p.urutan, p.status,
			p.created_at, p.updated_at
		FROM penugasan p
		LEFT JOIN distribusi_pelaksana dp ON p.distribusi_pelaksana_id = dp.id
		WHERE dp.programmer_id = $1
		ORDER BY p.urutan ASC, p.created_at ASC
	`, programmerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PenugasanItem
	for rows.Next() {
		var item PenugasanItem
		if err := rows.Scan(
			&item.ID, &item.DistribusiPelaksanaID,
			&item.Judul, &item.Deskripsi, &item.Deadline, &item.Prioritas,
			&item.EstimasiHari, &item.Urutan, &item.Status,
			&item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func fetchLaporanByProgrammerID(programmerID string) ([]LaporanItem, error) {
	rows, err := config.DB.Query(`
		SELECT
			l.id,
			u.id, u.username, u.full_name, u.profile_picture,
			l.laporan_progress, l.status,
			l.created_at, l.updated_at
		FROM laporan_kinerja l
		LEFT JOIN users u ON l.programmer_id = u.id
		WHERE l.programmer_id = $1
		ORDER BY l.created_at DESC
	`, programmerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []LaporanItem
	for rows.Next() {
		var item LaporanItem
		if err := rows.Scan(
			&item.ID,
			&item.Programmer.ID, &item.Programmer.Username, &item.Programmer.FullName, &item.Programmer.ProfilePicture,
			&item.LaporanProgress, &item.Status,
			&item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

// ── Verifikator ───────────────────────────────────────────────────────────────

func fetchLaporanPendingVerifikasi() ([]VerifikatorLaporanItem, error) {
	rows, err := config.DB.Query(`
		SELECT
			l.id,
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo,
			p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran, p.status,
			pu.id, pu.username, pu.full_name, pu.profile_picture,
			p.created_at, p.updated_at,
			u.id, u.username, u.full_name, u.profile_picture,
			l.laporan_progress, l.status,
			l.created_at, l.updated_at
		FROM laporan_kinerja l
		LEFT JOIN permintaan p ON l.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users pu ON p.created_by = pu.id
		LEFT JOIN users u ON l.programmer_id = u.id
		JOIN verifikasi v ON v.laporan_id = l.id
		WHERE v.status_verified = 'pending'
		ORDER BY l.created_at ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []VerifikatorLaporanItem
	for rows.Next() {
		var item VerifikatorLaporanItem
		var perm PermintaanItem
		if err := rows.Scan(
			&item.ID,
			&perm.ID, &perm.Pemda.ID, &perm.Pemda.Name, &perm.Pemda.Logo,
			&perm.Aplikasi.ID, &perm.Aplikasi.Name, &perm.Aplikasi.Logo,
			&perm.Menu, &perm.KondisiAwal, &perm.KondisiDiharapkan,
			&perm.TanggalPesanan, &perm.TanggalDeadline, &perm.Lampiran, &perm.Status,
			&perm.Pembuat.ID, &perm.Pembuat.Username, &perm.Pembuat.FullName, &perm.Pembuat.ProfilePicture,
			&perm.CreatedAt, &perm.UpdatedAt,
			&item.Programmer.ID, &item.Programmer.Username, &item.Programmer.FullName, &item.Programmer.ProfilePicture,
			&item.LaporanProgress, &item.Status,
			&item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		perm.Distribusi = []DistribusiItem{}
		perm.Laporan = []LaporanItem{}
		item.Permintaan = perm
		items = append(items, item)
	}
	return items, rows.Err()
}

// ── Activity helpers ──────────────────────────────────────────────────────────

func scanActivities(rows interface {
	Next() bool
	Scan(...interface{}) error
	Err() error
	Close() error
}) ([]ActivityItem, error) {
	defer rows.Close()
	var activities []ActivityItem
	for rows.Next() {
		var item ActivityItem
		var refID, actorID, actorUsername, actorFullName, actorPP string
		var pemdaName, aplikasiName, menu string
		if err := rows.Scan(
			&item.ActivityType, &item.ActivityID, &refID,
			&actorID, &actorUsername, &actorFullName, &actorPP,
			&pemdaName, &aplikasiName, &menu,
			&item.Status, &item.Komentar,
			&item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		item.RefID = refID
		item.Actor = ActorInfo{
			ID:             actorID,
			Username:       actorUsername,
			FullName:       actorFullName,
			ProfilePicture: actorPP,
		}
		item.PemdaName = pemdaName
		item.AplikasiName = aplikasiName
		item.Menu = menu
		activities = append(activities, item)
	}
	return activities, rows.Err()
}

// ── Activity — SuperAdmin ─────────────────────────────────────────────────────

func fetchActivitiesSuperAdmin() ([]ActivityItem, error) {
	rows, err := config.DB.Query(`
		SELECT activity_type, activity_id, ref_id,
			actor_id, actor_username, actor_full_name, actor_profile_picture,
			pemda_name, aplikasi_name, menu,
			status, komentar, created_at, updated_at
		FROM (

			SELECT 'permintaan' AS activity_type, p.id AS activity_id, p.id AS ref_id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), p.menu,
				p.status, '' AS komentar, p.created_at, p.updated_at
			FROM permintaan p
			LEFT JOIN users u ON p.created_by = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

			UNION ALL

			SELECT 'distribusi', d.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', COALESCE(d.komentar,''), d.created_at, d.updated_at
			FROM distribusi d
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON d.admin_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

			UNION ALL

			SELECT 'pelaksana', dp.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', '', dp.created_at, dp.updated_at
			FROM distribusi_pelaksana dp
			LEFT JOIN distribusi d ON dp.distribusi_id = d.id
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON dp.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

			UNION ALL

			SELECT 'penugasan', pg.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				pg.status, COALESCE(pg.judul,''), pg.created_at, pg.updated_at
			FROM penugasan pg
			LEFT JOIN distribusi_pelaksana dp ON pg.distribusi_pelaksana_id = dp.id
			LEFT JOIN distribusi d ON dp.distribusi_id = d.id
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON dp.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

			UNION ALL

			SELECT 'laporan', l.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				l.status, COALESCE(l.laporan_progress,''), l.created_at, l.updated_at
			FROM laporan_kinerja l
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON l.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

			UNION ALL

			SELECT 'verifikasi', v.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				v.status_verified, COALESCE(v.komentar,''), v.created_at, v.updated_at
			FROM verifikasi v
			LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON v.verifikator_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

			UNION ALL

			SELECT 'komentar_distribusi', kd.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', COALESCE(kd.komentars,''), kd.created_at, kd.updated_at
			FROM komentar_distribusi kd
			LEFT JOIN distribusi d ON kd.distribusi_id = d.id
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON kd.user_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

			UNION ALL

			SELECT 'komentar_laporan', kl.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', COALESCE(kl.komentar,''), kl.created_at, kl.updated_at
			FROM komentar_laporan kl
			LEFT JOIN laporan_kinerja l ON kl.laporan_id = l.id
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON kl.user_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

		) AS activities
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	return scanActivities(rows)
}

// ── Activity — Admin ──────────────────────────────────────────────────────────

func fetchActivitiesAdmin(adminID string) ([]ActivityItem, error) {
	rows, err := config.DB.Query(`
		SELECT activity_type, activity_id, ref_id,
			actor_id, actor_username, actor_full_name, actor_profile_picture,
			pemda_name, aplikasi_name, menu,
			status, komentar, created_at, updated_at
		FROM (

			SELECT 'distribusi', d.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', COALESCE(d.komentar,''), d.created_at, d.updated_at
			FROM distribusi d
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON d.admin_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE d.admin_id = $1

			UNION ALL

			SELECT 'pelaksana', dp.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', '', dp.created_at, dp.updated_at
			FROM distribusi_pelaksana dp
			LEFT JOIN distribusi d ON dp.distribusi_id = d.id
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON dp.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE d.admin_id = $1

			UNION ALL

			SELECT 'penugasan', pg.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				pg.status, COALESCE(pg.judul,''), pg.created_at, pg.updated_at
			FROM penugasan pg
			LEFT JOIN distribusi_pelaksana dp ON pg.distribusi_pelaksana_id = dp.id
			LEFT JOIN distribusi d ON dp.distribusi_id = d.id
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON dp.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE d.admin_id = $1

			UNION ALL

			SELECT 'laporan', l.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				l.status, COALESCE(l.laporan_progress,''), l.created_at, l.updated_at
			FROM laporan_kinerja l
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON l.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE p.id IN (SELECT permintaan_id FROM distribusi WHERE admin_id = $1)

			UNION ALL

			SELECT 'verifikasi', v.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				v.status_verified, COALESCE(v.komentar,''), v.created_at, v.updated_at
			FROM verifikasi v
			LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON v.verifikator_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE p.id IN (SELECT permintaan_id FROM distribusi WHERE admin_id = $1)

			UNION ALL

			SELECT 'komentar_distribusi', kd.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', COALESCE(kd.komentars,''), kd.created_at, kd.updated_at
			FROM komentar_distribusi kd
			LEFT JOIN distribusi d ON kd.distribusi_id = d.id
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON kd.user_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE d.admin_id = $1

			UNION ALL

			SELECT 'penilaian', pn.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				pn.ketepatan_waktu, COALESCE(pn.komentar,''), pn.created_at, pn.updated_at
			FROM penilaian pn
			LEFT JOIN distribusi d ON pn.distribusi_id = d.id
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON pn.penilai_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE pn.penilai_id = $1

		) AS activities
		ORDER BY created_at DESC
	`, adminID)
	if err != nil {
		return nil, err
	}
	return scanActivities(rows)
}

// ── Activity — Programmer ─────────────────────────────────────────────────────

func fetchActivitiesProgrammer(programmerID string) ([]ActivityItem, error) {
	rows, err := config.DB.Query(`
		SELECT activity_type, activity_id, ref_id,
			actor_id, actor_username, actor_full_name, actor_profile_picture,
			pemda_name, aplikasi_name, menu,
			status, komentar, created_at, updated_at
		FROM (

			SELECT 'pelaksana', dp.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', '', dp.created_at, dp.updated_at
			FROM distribusi_pelaksana dp
			LEFT JOIN distribusi d ON dp.distribusi_id = d.id
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON dp.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE dp.programmer_id = $1

			UNION ALL

			SELECT 'penugasan', pg.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				pg.status, COALESCE(pg.judul,''), pg.created_at, pg.updated_at
			FROM penugasan pg
			LEFT JOIN distribusi_pelaksana dp ON pg.distribusi_pelaksana_id = dp.id
			LEFT JOIN distribusi d ON dp.distribusi_id = d.id
			LEFT JOIN permintaan p ON d.permintaan_id = p.id
			LEFT JOIN users u ON dp.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE dp.programmer_id = $1

			UNION ALL

			SELECT 'laporan', l.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				l.status, COALESCE(l.laporan_progress,''), l.created_at, l.updated_at
			FROM laporan_kinerja l
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON l.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE l.programmer_id = $1

			UNION ALL

			SELECT 'verifikasi', v.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				v.status_verified, COALESCE(v.komentar,''), v.created_at, v.updated_at
			FROM verifikasi v
			LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON v.verifikator_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE l.programmer_id = $1

			UNION ALL

			SELECT 'komentar_laporan', kl.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', COALESCE(kl.komentar,''), kl.created_at, kl.updated_at
			FROM komentar_laporan kl
			LEFT JOIN laporan_kinerja l ON kl.laporan_id = l.id
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON kl.user_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE l.programmer_id = $1

		) AS activities
		ORDER BY created_at DESC
	`, programmerID)
	if err != nil {
		return nil, err
	}
	return scanActivities(rows)
}

// ── Activity — Verifikator ────────────────────────────────────────────────────

func fetchActivitiesVerifikator(verifikatorID string) ([]ActivityItem, error) {
	rows, err := config.DB.Query(`
		SELECT activity_type, activity_id, ref_id,
			actor_id, actor_username, actor_full_name, actor_profile_picture,
			pemda_name, aplikasi_name, menu,
			status, komentar, created_at, updated_at
		FROM (

			SELECT 'laporan', l.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				l.status, COALESCE(l.laporan_progress,''), l.created_at, l.updated_at
			FROM laporan_kinerja l
			JOIN verifikasi v ON v.laporan_id = l.id AND v.is_submitted_to_verified = true
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON l.programmer_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

			UNION ALL

			SELECT 'verifikasi', v.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				v.status_verified, COALESCE(v.komentar,''), v.created_at, v.updated_at
			FROM verifikasi v
			LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON v.verifikator_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE v.verifikator_id = $1

			UNION ALL

			SELECT 'komentar_laporan', kl.id, p.id,
				COALESCE(u.id,''), COALESCE(u.username,''), COALESCE(u.full_name,''), COALESCE(u.profile_picture,''),
				COALESCE(mp.name,''), COALESCE(ma.name,''), COALESCE(p.menu,''),
				'', COALESCE(kl.komentar,''), kl.created_at, kl.updated_at
			FROM komentar_laporan kl
			LEFT JOIN laporan_kinerja l ON kl.laporan_id = l.id
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN users u ON kl.user_id = u.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			WHERE EXISTS (
				SELECT 1 FROM verifikasi v2
				WHERE v2.laporan_id = l.id AND v2.verifikator_id = $1
			)

		) AS activities
		ORDER BY created_at DESC
	`, verifikatorID)
	if err != nil {
		return nil, err
	}
	return scanActivities(rows)
}
