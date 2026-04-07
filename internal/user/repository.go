package user

import (
	"aplikasi-internal/config"
	"fmt"
	"strings"
)

func GetAll() ([]UserResponseDetail, error) {
	rows, err := config.DB.Query(
		`SELECT u.id, u.role_id, r.name, r.description, u.username, u.full_name, u.profile_picture, u.is_active, u.created_at, u.updated_at FROM users u
		LEFT JOIN roles r ON u.role_id = r.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []UserResponseDetail

	for rows.Next() {
		var data UserResponseDetail
		err := rows.Scan(&data.ID, &data.Role.ID, &data.Role.Name, &data.Role.Description, &data.Username, &data.FullName, &data.ProfilePicture, &data.IsActive, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}

	return users, err
}

func GetId(id string) (UserResponseDetail, error) {
	var data UserResponseDetail
	err := config.DB.QueryRow(`SELECT u.id, u.role_id, r.name, r.description, u.username, u.full_name, u.profile_picture, u.is_active, u.created_at, u.updated_at FROM users u
	LEFT JOIN roles r ON u.role_id = r.id
	WHERE u.id=$1`, id).
		Scan(&data.ID, &data.Role.ID, &data.Role.Name, &data.Role.Description, &data.Username, &data.FullName, &data.ProfilePicture, &data.IsActive, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return UserResponseDetail{}, err
	}

	return data, err
}

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
		INSERT INTO users (role_id, username, full_name, password, profile_picture)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, is_active, created_at, updated_at
	`

	return config.DB.QueryRow(
		query,
		user.RoleID,
		user.Username,
		user.FullName,
		hashedPassword,
		user.ProfilePicture,
	).Scan(
		&user.ID,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

func UpdateUser(id string, fields map[string]interface{}) error {
	setClauses := []string{}
	args := []interface{}{}
	i := 1
	for col, val := range fields {
		setClauses = append(setClauses, col+" = $"+fmt.Sprintf("%d", i))
		args = append(args, val)
		i++
	}
	args = append(args, id)
	query := "UPDATE users SET " + strings.Join(setClauses, ", ") + ", updated_at = NOW() WHERE id = $" + fmt.Sprintf("%d", i)
	_, err := config.DB.Exec(query, args...)
	return err
}

func UpdateProfilePicture(id string, profilePicture string) error {
	query := `UPDATE users SET profile_picture = $1, updated_at = NOW() WHERE id = $2`
	_, err := config.DB.Exec(query, profilePicture, id)
	return err
}

func GetUserByUsername(username string) (*UserRole, error) {

	query := `
	SELECT u.id, u.role_id, r.name, u.username, u.full_name, u.password, u.is_active
	FROM users u
	LEFT JOIN roles r ON u.role_id = r.id 
	WHERE username = $1
	`

	row := config.DB.QueryRow(query, username)

	var user UserRole

	err := row.Scan(
		&user.ID,
		&user.RoleID,
		&user.RoleName,
		&user.Username,
		&user.FullName,
		&user.Password,
		&user.IsActive,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeactivateUser(id string) error {

	query := `
	UPDATE users
	SET is_active = false,
	    updated_at = NOW()
	WHERE id = $1
	`

	_, err := config.DB.Exec(query, id)

	return err
}
