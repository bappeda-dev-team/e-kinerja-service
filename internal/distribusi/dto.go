package distribusi

import "time"

// === Request ===

type DistribusiRequest struct {
	PermintaanID string `json:"permintaan_id" validate:"required,uuid4"`
	Komentar     string `json:"komentar"`
}

// === Response ===

type DistribusiResponse struct {
	ID           string    `json:"id"`
	PermintaanID string    `json:"permintaan_id"`
	AdminID      string    `json:"admin_id"`
	Komentar     string    `json:"komentar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DistribusiDetailResponse struct {
	ID        string    `json:"id"`
	Pemda     string    `json:"pemda"`
	Aplikasi  string    `json:"aplikasi"`
	Admin     string    `json:"admin"`
	Komentar  string    `json:"komentar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
