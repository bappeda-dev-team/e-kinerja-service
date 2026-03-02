package pelaksana

import "time"

type Pelaksana struct {
	ID           string    `json:"id"`
	DistribusiID string    `json:"distribusi_id"`
	ProgrammerID string    `json:"programmer_id"`
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