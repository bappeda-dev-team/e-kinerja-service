package master_aplikasi

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAllMasterAplikasi() ([]MasterAplikasi, error) {
	rows, err := config.DB.Query("SELECT id, name, logo, created_at, updated_at FROM master_aplikasi")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var master_aplikasi []MasterAplikasi

	for rows.Next() {
		var aplikasi MasterAplikasi
		err := rows.Scan(&aplikasi.ID, &aplikasi.Name, &aplikasi.Logo, &aplikasi.CreatedAt, &aplikasi.UpdatedAt)
		if err != nil {
			return nil, err
		}
		master_aplikasi = append(master_aplikasi, aplikasi)
	}

	return master_aplikasi, err

}

func GetMasterAplikasiId(id string) (MasterAplikasi, error) {
	var aplikasi MasterAplikasi
	err := config.DB.QueryRow("SELECT id, name, logo, created_at, updated_at FROM master_aplikasi WHERE id=$1", id).
		Scan(&aplikasi.ID, &aplikasi.Name, &aplikasi.Logo, &aplikasi.CreatedAt, &aplikasi.UpdatedAt)

	if err != nil {
		return MasterAplikasi{}, err
	}

	return aplikasi, err
}

func CreateMasterAplikasi(data *MasterAplikasi) error {
	query := `
		INSERT INTO master_aplikasi (name, logo)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.Name,
		data.Logo,
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
		    logo = CASE WHEN $2 = '' THEN logo ELSE $2 END,
		    updated_at = NOW()
		WHERE id = $3
		RETURNING updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.Name,
		data.Logo,
		id,
	).Scan(&data.UpdatedAt)

	return err
}

func UpdateLogoMasterAplikasi(id string, logoURL string) error {
	result, err := config.DB.Exec(`UPDATE master_aplikasi SET logo = $1, updated_at = NOW() WHERE id = $2`, logoURL, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
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


