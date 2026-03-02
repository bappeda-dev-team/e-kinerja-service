package distribusi

import "time"

type Distribusi struct {
	ID           string    `json:"id"`
	Pemda 		 string    `json:"pemda"`
	Aplikasi 	 string    `json:"aplikasi"`
	Admin 	     string    `json:"admin"`
	Komentar 	 string    `json:"komentar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type APIResponse struct {
	Code    int         `json:"code"`
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Errors  interface{} `json:"errors,omitempty"`
}