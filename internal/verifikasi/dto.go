package verifikasi

import "time"

// === Request ===

type VerifikasiRequest struct {
	LaporanID      string `json:"laporan_id" validate:"required,uuid4"`
	Komentar       string `json:"komentar"`
	StatusVerified string `json:"status_verified" validate:"required,oneof=pending approved revision"`
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

type ProgrammerInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type LaporanInfo struct {
	ID              string        `json:"id"`
	LaporanProgress string        `json:"laporan_progress"`
	Status          string        `json:"status"`
	Programmer      ProgrammerInfo `json:"programmer"`
}

type VerifikatorInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type VerifikasiDetailResponse struct {
	ID             string          `json:"id"`
	Laporan        LaporanInfo     `json:"laporan"`
	Verifikator    VerifikatorInfo `json:"verifikator"`
	Komentar       string          `json:"komentar"`
	StatusVerified string          `json:"status_verified"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
