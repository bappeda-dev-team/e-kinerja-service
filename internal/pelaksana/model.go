package pelaksana

import "time"

type Pelaksana struct {
	ID           string    `json:"id"`
	DistribusiID string    `json:"distribusi_id"`
	ProgrammerID string    `json:"programmer_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type PelaksanaNama struct {
	ID           string    `json:"id"`
	Pemda		 string    `json:"pemda"`
	Aplikasi	 string    `json:"aplikasi"`
	Programmer 	 string    `json:"programmer"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PelaksanaRequest struct {
	DistribusiID string `json:"distribusi_id" binding:"required"`
	ProgrammerID string `json:"programmer_id" binding:"required"`
}