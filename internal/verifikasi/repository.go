package verifikasi

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAll() ([]VerifikasiResponse, error) {
	rows, err := config.DB.Query(
		`SELECT id, laporan_id, verifikator_id, komentar, status_verified, created_at, updated_at FROM verifikasi`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var verifikasi []VerifikasiResponse

	for rows.Next() {
		var data VerifikasiResponse
		err := rows.Scan(&data.ID, &data.LaporanID, &data.VerifikatorID, &data.Komentar, &data.StatusVerified, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		verifikasi = append(verifikasi, data)
	}

	return verifikasi, err

}

func GetId(id string) (VerifikasiResponse, error) {
	var data VerifikasiResponse
	err := config.DB.QueryRow(`SELECT id, laporan_id, verifikator_id, komentar, status_verified, created_at, updated_at FROM verifikasi WHERE id=$1`, id).
		Scan(&data.ID, &data.LaporanID, &data.VerifikatorID, &data.Komentar, &data.StatusVerified, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return VerifikasiResponse{}, err
	}

	return data, err
}

func GetAllDetail() ([]VerifikasiDetailResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			v.id,
			COALESCE(d.id::text, ''),
			l.id, l.laporan_progress, l.status,
			up.id, up.username, up.full_name, up.profile_picture,
			uv.id, uv.username, uv.full_name, uv.profile_picture,
			v.komentar, v.status_verified,
			v.created_at, v.updated_at
		FROM verifikasi v
		LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
		LEFT JOIN distribusi d ON d.permintaan_id = l.permintaan_id
		LEFT JOIN users up ON l.programmer_id = up.id
		LEFT JOIN users uv ON v.verifikator_id = uv.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var verifikasi []VerifikasiDetailResponse
	for rows.Next() {
		var data VerifikasiDetailResponse
		err := rows.Scan(
			&data.ID,
			&data.DistribusiID,
			&data.Laporan.ID, &data.Laporan.LaporanProgress, &data.Laporan.Status,
			&data.Laporan.Programmer.ID, &data.Laporan.Programmer.Username, &data.Laporan.Programmer.FullName, &data.Laporan.Programmer.ProfilePicture,
			&data.Verifikator.ID, &data.Verifikator.Username, &data.Verifikator.FullName, &data.Verifikator.ProfilePicture,
			&data.Komentar, &data.StatusVerified,
			&data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		verifikasi = append(verifikasi, data)
	}
	return verifikasi, err
}
func GetAllBylaporan() ([]VerifikasiDetailResponse, error) {
	rows, err := config.DB.Query(`
		SELECT DISTINCT ON (l.id)
			v.id,
			COALESCE(d.id::text, ''),
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			l.id, l.laporan_progress, l.status,
			up.id, up.username, up.full_name, up.profile_picture,
			uv.id, uv.username, uv.full_name, uv.profile_picture,
			v.komentar, v.status_verified,
			v.created_at, v.updated_at
		FROM verifikasi v
		LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
		LEFT JOIN permintaan p ON l.permintaan_id = p.id
		LEFT JOIN distribusi d ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users up ON l.programmer_id = up.id
		LEFT JOIN users uv ON v.verifikator_id = uv.id
		ORDER BY l.id, v.updated_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var verifikasi []VerifikasiDetailResponse
	for rows.Next() {
		var data VerifikasiDetailResponse
		err := rows.Scan(
			&data.ID,
			&data.DistribusiID,
			&data.Permintaan.ID, &data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo,
			&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
			&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
			&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
			&data.Laporan.ID, &data.Laporan.LaporanProgress, &data.Laporan.Status,
			&data.Laporan.Programmer.ID, &data.Laporan.Programmer.Username, &data.Laporan.Programmer.FullName, &data.Laporan.Programmer.ProfilePicture,
			&data.Verifikator.ID, &data.Verifikator.Username, &data.Verifikator.FullName, &data.Verifikator.ProfilePicture,
			&data.Komentar, &data.StatusVerified,
			&data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		verifikasi = append(verifikasi, data)
	}
	return verifikasi, err
}

func GetByIdDetail(id string) (VerifikasiDetailResponse, error) {
	var data VerifikasiDetailResponse
	err := config.DB.QueryRow(`
		SELECT
			v.id,
			COALESCE(d.id::text, ''),
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			l.id, l.laporan_progress, l.status,
			up.id, up.username, up.full_name, up.profile_picture,
			uv.id, uv.username, uv.full_name, uv.profile_picture,
			v.komentar, v.status_verified,
			v.created_at, v.updated_at
		FROM verifikasi v
		LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
		LEFT JOIN permintaan p ON l.permintaan_id = p.id
		LEFT JOIN distribusi d ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users up ON l.programmer_id = up.id
		LEFT JOIN users uv ON v.verifikator_id = uv.id
		WHERE v.id = $1
	`, id).Scan(
		&data.ID,
		&data.DistribusiID,
		&data.Permintaan.ID, &data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo,
		&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
		&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
		&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
		&data.Laporan.ID, &data.Laporan.LaporanProgress, &data.Laporan.Status,
		&data.Laporan.Programmer.ID, &data.Laporan.Programmer.Username, &data.Laporan.Programmer.FullName, &data.Laporan.Programmer.ProfilePicture,
		&data.Verifikator.ID, &data.Verifikator.Username, &data.Verifikator.FullName, &data.Verifikator.ProfilePicture,
		&data.Komentar, &data.StatusVerified,
		&data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return VerifikasiDetailResponse{}, err
	}
	return data, err
}

func Create(data *Verifikasi) error {
	return config.DB.QueryRow(`
		INSERT INTO verifikasi (laporan_id, verifikator_id, komentar, status_verified)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`, data.LaporanID, data.VerifikatorID, data.Komentar, data.StatusVerified).Scan(
		&data.ID, &data.CreatedAt, &data.UpdatedAt,
	)
}

func Update(id string, data *Verifikasi) error {
	return config.DB.QueryRow(`
		UPDATE verifikasi
		SET laporan_id = $1,
			verifikator_id = $2,
			komentar = $3,
			status_verified = $4,
		    updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`, data.LaporanID, data.VerifikatorID, data.Komentar, data.StatusVerified, id).Scan(&data.UpdatedAt)
}

func Delete(id string) error {
	query := `
		DELETE FROM verifikasi
		WHERE id = $1
	`

	result, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
