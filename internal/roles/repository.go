package roles

import (
	"aplikasi-internal/config"
)

func GetAllRoles() ([]Roles, error){
	rows, err := config.DB.Query("SELECT id, name, description, created_at, updated_at FROM roles")
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