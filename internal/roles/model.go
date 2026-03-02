package roles

import "time"

type Roles struct {
	ID        	 string    `json:"id"`
	Name    	 string    `json:"name"`
	Description  string    `json:"description"`
	CreatedAt 	 time.Time `json:"created_at"`
	UpdatedAt 	 time.Time `json:"updated_at"`
}

type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Errors  interface{} `json:"errors,omitempty"`
}