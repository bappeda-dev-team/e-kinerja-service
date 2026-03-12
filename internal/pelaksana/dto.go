package pelaksana

import "time"

// === Request ===

type PelaksanaRequest struct {
	DistribusiID string `json:"distribusi_id" validate:"required,uuid4"`
	ProgrammerID string `json:"programmer_id" validate:"required,uuid4"`
}

// === Response ===

type PelaksanaResponse struct {
	ID           string    `json:"id"`
	DistribusiID string    `json:"distribusi_id"`
	ProgrammerID string    `json:"programmer_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PelaksanaDetailResponse struct {
	ID         string    `json:"id"`
	Pemda      string    `json:"pemda"`
	Aplikasi   string    `json:"aplikasi"`
	Programmer string    `json:"programmer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
