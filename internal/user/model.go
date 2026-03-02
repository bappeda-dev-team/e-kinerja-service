package user

import "time"

type User struct {
	ID        string    `json:"id"`
	RoleID    string    `json:"role_id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"-"` // jangan tampilkan di JSON
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Roles struct {
	ID        	 string    `json:"id"`
	Name    	 string    `json:"name"`
}

type APIResponse struct {
	Code    int         `json:"code"`
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Errors  interface{} `json:"errors,omitempty"`
}