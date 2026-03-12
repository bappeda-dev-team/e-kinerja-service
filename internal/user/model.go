package user

import "time"

type User struct {
	ID             string    `json:"id"`
	RoleID         string    `json:"role_id"`
	Username       string    `json:"username"`
	FullName       string    `json:"full_name"`
	Password       string    `json:"-"`
	ProfilePicture string    `json:"profile_picture"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// dipakai internal untuk login (butuh Password & RoleName)
type UserRole struct {
	ID       string `json:"id"`
	RoleID   string `json:"role_id"`
	RoleName string `json:"name"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Password string `json:"-"`
	IsActive bool   `json:"is_active"`
}