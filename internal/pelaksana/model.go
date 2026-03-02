package pelaksana

import "time"

type Pelaksana struct {
	ID           string    `json:"id"`
	Pemda		 string    `json:"pemda"`
	Aplikasi	 string    `json:"aplikasi"`
	Programmer 	 string    `json:"programmer"`
	CreatedAt    time.Time `json:"created_at"`
}

type APIResponse struct {
	Code    int         `json:"code"`
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Errors  interface{} `json:"errors,omitempty"`
}