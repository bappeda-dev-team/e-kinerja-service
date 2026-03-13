package user

import "time"

// === Request ===

type RegisterRequest struct {
	RoleID   string `json:"role_id"   form:"role_id"   validate:"required"`
	Username string `json:"username"  form:"username"  validate:"required,min=4"`
	FullName string `json:"full_name" form:"full_name" validate:"required,min=3"`
	Password string `json:"password"  form:"password"  validate:"required,min=6"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// === Response ===

type UserResponse struct {
	ID             string    `json:"id"`
	RoleID         string    `json:"role_id"`
	Username       string    `json:"username"`
	FullName       string    `json:"full_name"`
	ProfilePicture string    `json:"profile_picture"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
