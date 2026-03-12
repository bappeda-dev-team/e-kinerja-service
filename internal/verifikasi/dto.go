package verifikasi

import "time"

// === Request ===

type VerifikasiRequest struct {
	LaporanID      string `json:"laporan_id" validate:"required,uuid4"`
	Komentar       string `json:"komentar"`
	StatusVerified string `json:"status_verified" validate:"required"`
}

// === Response ===

type VerifikasiResponse struct {
	ID             string    `json:"id"`
	LaporanID      string    `json:"laporan_id"`
	VerifikatorID  string    `json:"verifikator_id"`
	Komentar       string    `json:"komentar"`
	StatusVerified string    `json:"status_verified"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
