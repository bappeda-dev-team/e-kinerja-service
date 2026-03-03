package master_aplikasi

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAllMasterAplikasi() ([]MasterAplikasi, error) {
	rows, err := config.DB.Query("SELECT id, name, created_at, updated_at FROM master_aplikasi")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var master_aplikasi []MasterAplikasi

	for rows.Next() {
		var aplikasi MasterAplikasi
		err := rows.Scan(&aplikasi.ID, &aplikasi.Name, &aplikasi.CreatedAt, &aplikasi.UpdatedAt)
		if err != nil {
			return nil, err
		}
		master_aplikasi = append(master_aplikasi, aplikasi)
	}

	return master_aplikasi, err

}

func GetMasterAplikasiId(id string) (MasterAplikasi, error){
	var aplikasi MasterAplikasi
	err := config.DB.QueryRow("SELECT id, name, created_at, updated_at FROM master_aplikasi WHERE id=$1", id).
		Scan(&aplikasi.ID, &aplikasi.Name, &aplikasi.CreatedAt, &aplikasi.UpdatedAt)

	if err != nil {
		return MasterAplikasi{}, err
	}

	return aplikasi, err
}

func CreateMasterAplikasi(data *MasterAplikasi) error {
	query := `
		INSERT INTO master_aplikasi (name)
		VALUES ($1)
		RETURNING id, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.Name,
	).Scan(
		&data.ID,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return err
}

func UpdateMasterAplikasi(id string, data *MasterAplikasi) error {
	query := `
		UPDATE master_aplikasi
		SET name = $1,
		    updated_at = NOW()
		WHERE id = $2
		RETURNING updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.Name,
		id,
	).Scan(&data.UpdatedAt)

	return err
}

func DeleteMasterAplikasi(id string) error {
	query := `
		DELETE FROM master_aplikasi
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


