package verifikasi

import "time"

type Verifikasi struct {
	ID         	   string    `json:"id"`
	LaporanID 	   string    `json:"laporan_id"`
	VerifikatorID  string    `json:"verifikator_id"`
	Komentar  	   string    `json:"komentar"`
	StatusVerified string    `json:"status_verified"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type VerifikasiRequest struct {
	LaporanID 		string `json:"laporan_id" binding:"required"`
	VerifikatorID	string `json:"verifikator_id" binding:"required"`
	Komentar  	    string `json:"komentar"`
	StatusVerified  string `json:"status_verified" binding:"required"`
}