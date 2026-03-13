package master_pemda

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAllMasterPemda() ([]MasterPemda, error) {
	rows, err := config.DB.Query("SELECT id, name, logo, created_at, updated_at FROM master_pemda")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var master_pemda []MasterPemda

	for rows.Next() {
		var pemda MasterPemda
		err := rows.Scan(&pemda.ID, &pemda.Name, &pemda.Logo, &pemda.CreatedAt, &pemda.UpdatedAt)
		if err != nil {
			return nil, err
		}
		master_pemda = append(master_pemda, pemda)
	}

	return master_pemda, err

}

func GetMasterPemdaId(id string) (MasterPemda, error) {
	var pemda MasterPemda
	err := config.DB.QueryRow("SELECT id, name, logo, created_at, updated_at FROM master_pemda WHERE id=$1", id).
		Scan(&pemda.ID, &pemda.Name, &pemda.Logo, &pemda.CreatedAt, &pemda.UpdatedAt)

	if err != nil {
		return MasterPemda{}, err
	}

	return pemda, err
}

func CreateMasterPemda(data *MasterPemda) error {
	query := `
		INSERT INTO master_pemda (name, logo)
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

func UpdateMasterPemda(id string, data *MasterPemda) error {
	query := `
		UPDATE master_pemda
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

func UpdateLogoMasterPemda(id string, logoURL string) error {
	result, err := config.DB.Exec(`UPDATE master_pemda SET logo = $1, updated_at = NOW() WHERE id = $2`, logoURL, id)
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

func DeleteMasterPemda(id string) error {
	query := `
		DELETE FROM master_pemda
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