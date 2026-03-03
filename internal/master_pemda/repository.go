package master_pemda

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAllMasterPemda() ([]MasterPemda, error) {
	rows, err := config.DB.Query("SELECT id, name, created_at, updated_at FROM master_pemda")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var master_pemda []MasterPemda

	for rows.Next() {
		var pemda MasterPemda
		err := rows.Scan(&pemda.ID, &pemda.Name, &pemda.CreatedAt, &pemda.UpdatedAt)
		if err != nil {
			return nil, err
		}
		master_pemda = append(master_pemda, pemda)
	}

	return master_pemda, err

}

func GetMasterPemdaId(id string) (MasterPemda, error){
	var pemda MasterPemda
	err := config.DB.QueryRow("SELECT id, name, created_at, updated_at FROM master_pemda WHERE id=$1", id).
		Scan(&pemda.ID, &pemda.Name, &pemda.CreatedAt, &pemda.UpdatedAt)

	if err != nil {
		return MasterPemda{}, err
	}

	return pemda, err
}

func CreateMasterPemda(data *MasterPemda) error {
	query := `
		INSERT INTO master_pemda (name)
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

func UpdateMasterPemda(id string, data *MasterPemda) error {
	query := `
		UPDATE master_pemda
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