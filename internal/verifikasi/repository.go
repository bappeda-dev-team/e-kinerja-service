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
			l.id, l.laporan_progress, l.status,
			up.id, up.username, up.full_name,
			uv.id, uv.username, uv.full_name,
			v.komentar, v.status_verified,
			v.created_at, v.updated_at
		FROM verifikasi v
		LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
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
			&data.Laporan.ID, &data.Laporan.LaporanProgress, &data.Laporan.Status,
			&data.Laporan.Programmer.ID, &data.Laporan.Programmer.Username, &data.Laporan.Programmer.FullName,
			&data.Verifikator.ID, &data.Verifikator.Username, &data.Verifikator.FullName,
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
			l.id, l.laporan_progress, l.status,
			up.id, up.username, up.full_name,
			uv.id, uv.username, uv.full_name,
			v.komentar, v.status_verified,
			v.created_at, v.updated_at
		FROM verifikasi v
		LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
		LEFT JOIN users up ON l.programmer_id = up.id
		LEFT JOIN users uv ON v.verifikator_id = uv.id
		WHERE v.id = $1
	`, id).Scan(
		&data.ID,
		&data.Laporan.ID, &data.Laporan.LaporanProgress, &data.Laporan.Status,
		&data.Laporan.Programmer.ID, &data.Laporan.Programmer.Username, &data.Laporan.Programmer.FullName,
		&data.Verifikator.ID, &data.Verifikator.Username, &data.Verifikator.FullName,
		&data.Komentar, &data.StatusVerified,
		&data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return VerifikasiDetailResponse{}, err
	}
	return data, err
}

func Create(data *Verifikasi) error {
	query := `
		INSERT INTO verifikasi (laporan_id, verifikator_id, komentar, status_verified)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.LaporanID,
		data.VerifikatorID,
		data.Komentar,
		data.StatusVerified,
	).Scan(
		&data.ID,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return err
}

func Update(id string, data *Verifikasi) error {
	query := `
		UPDATE verifikasi
		SET laporan_id = $1,
			verifikator_id = $2,
			komentar = $3,
			status_verified = $4,
		    updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.LaporanID,
		data.VerifikatorID,
		data.Komentar,
		data.StatusVerified,
		id,
	).Scan(&data.UpdatedAt)

	return err
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