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