package master_aplikasi

import "time"

type MasterAplikasi struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type APIResponse struct {
	Code    int         `json:"code"`
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Errors  interface{} `json:"errors,omitempty"`
}

type CreateMasterAplikasiRequest struct {
	Name string `json:"name" binding:"required,min=3"`
}