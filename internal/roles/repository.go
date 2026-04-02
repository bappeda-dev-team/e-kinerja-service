package roles

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAllRoles() ([]Roles, error) {
	rows, err := config.DB.Query(`SELECT id, name, description, created_at, updated_at FROM roles`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Roles

	for rows.Next() {
		var role Roles
		err := rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, err
}

func GetId(id string) (Roles, error) {
	var data Roles
	err := config.DB.QueryRow(`SELECT id, name, description, created_at, updated_at FROM roles WHERE id=$1`, id).
		Scan(&data.ID, &data.Name, &data.Description, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return Roles{}, err
	}

	return data, err
}

func Update(id string, name string, description string) (Roles, error) {
	var data Roles
	err := config.DB.QueryRow(`
		UPDATE roles SET name = $1, description = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING id, name, description, created_at, updated_at
	`, name, description, id).
		Scan(&data.ID, &data.Name, &data.Description, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return Roles{}, err
	}

	return data, nil
}

func Delete(id string) error {
	result, err := config.DB.Exec(`DELETE FROM roles WHERE id = $1`, id)
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
