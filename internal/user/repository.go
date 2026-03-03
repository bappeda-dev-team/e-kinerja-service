package user

import "aplikasi-internal/config"


func IsUsernameExists(username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)`
	err := config.DB.QueryRow(query, username).Scan(&exists)
	return exists, err
}

func IsRoleExists(roleID string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM roles WHERE id = $1)`
	err := config.DB.QueryRow(query, roleID).Scan(&exists)
	return exists, err
}

func CreateUser(user *User, hashedPassword string) error {
	query := `
		INSERT INTO users (role_id, username, full_name, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, is_active, created_at, updated_at
	`

	return config.DB.QueryRow(
		query,
		user.RoleID,
		user.Username,
		user.FullName,
		hashedPassword,
	).Scan(
		&user.ID,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}
