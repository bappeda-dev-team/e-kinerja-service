package distribusi

import "time"

type Distribusi struct {
	ID           string    `json:"id"`
	PermintaanID string    `json:"permintaan_id"`
	AdminID 	 string    `json:"admin_id"`
	Komentar 	 string    `json:"komentar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DistribusiByNama struct {
	ID           string    `json:"id"`
	Pemda 		 string    `json:"pemda"`
	Aplikasi 	 string    `json:"aplikasi"`
	Admin 	     string    `json:"admin"`
	Komentar 	 string    `json:"komentar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DistribusiRequest struct {
	Komentar 		  string 	`json:"komentar"`
}