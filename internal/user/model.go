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


type APIResponse struct {
	Code    int         `json:"code"`
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Errors  interface{} `json:"errors,omitempty"`
}

type RegisterRequest struct {
	RoleID   string `json:"role_id" binding:"required"`
	Username string `json:"username" binding:"required,min=4"`
	FullName string `json:"full_name" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
}