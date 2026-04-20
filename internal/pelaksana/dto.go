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
	IsRead       bool      `json:"is_read"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DistribusiInfo struct {
	ID       	 string `json:"id"`
	PermintaanID string `json:"permintaan_id"`
	Pemda    	 string `json:"pemda"`
	Aplikasi 	 string `json:"aplikasi"`
	Komentar 	 string `json:"komentar"`
}

type ProgrammerInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type PelaksanaDetailResponse struct {
	ID         string         `json:"id"`
	Distribusi DistribusiInfo `json:"distribusi"`
	Programmer ProgrammerInfo `json:"programmer"`
	IsRead     bool           `json:"is_read"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}
