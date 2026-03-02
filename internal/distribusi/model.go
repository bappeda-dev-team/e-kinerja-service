package distribusi

import "time"

type Distribusi struct {
	ID           string    `json:"id"`
	PermintaanID string    `json:"permintaan_id"`
	AdminID 	 string    `json:"admin_id"`
	Komentar 	 string    `json:"komentar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Errors  interface{} `json:"errors,omitempty"`
}