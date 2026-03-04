package roles

import (
	"aplikasi-internal/config"
)

func GetAllRoles() ([]Roles, error){
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