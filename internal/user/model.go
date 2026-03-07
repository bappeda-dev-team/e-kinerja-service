package user

import "time"

type User struct {
	ID        string    `json:"id"`
	RoleID    string    `json:"role_id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	RoleID   string `json:"role_id" validate:"required"`
	Username string `json:"username" validate:"required,min=4"`
	FullName string `json:"full_name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
}