package laporan

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAll() ([]LaporanResponse, error) {
	rows, err := config.DB.Query(
		`SELECT id, permintaan_id, Programmer_id, laporan_progress, created_at, updated_at FROM laporan_kinerja`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanResponse

	for rows.Next() {
		var data LaporanResponse
		err := rows.Scan(&data.ID, &data.PermintaanID, &data.ProgrammerID, &data.LaporanProgress, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		laporan = append(laporan, data)
	}

	return laporan, err

}

func GetId(id string) (LaporanResponse, error) {
	var data LaporanResponse
	err := config.DB.QueryRow(`SELECT id, permintaan_id, Programmer_id, laporan_progress, created_at, updated_at FROM laporan_kinerja WHERE id=$1`, id).
		Scan(&data.ID, &data.PermintaanID, &data.ProgrammerID, &data.LaporanProgress, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return LaporanResponse{}, err
	}

	return data, err
}

func Create(data *Laporan) error {
	query := `
		INSERT INTO laporan_kinerja (permintaan_id, programmer_id, laporan_progress)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PermintaanID,
		data.ProgrammerID,
		data.LaporanProgress,
	).Scan(
		&data.ID,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return err
}

func Update(id string, data *Laporan) error {
	query := `
		UPDATE laporan_kinerja
		SET permintaan_id = $1,
			programmer_id = $2,
			laporan_progress = $3,
		    updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PermintaanID,
		data.ProgrammerID,
		data.LaporanProgress,
		id,
	).Scan(&data.UpdatedAt)

	return err
}

func Delete(id string) error {
	query := `
		DELETE FROM laporan_kinerja
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