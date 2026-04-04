package verifikasi

import (
	"aplikasi-internal/internal/permintaan"
	"time"
)

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

type PemdaInfo struct {
	ID			string     `json:"id"`
	Name		string	   `json:"name"`
	Logo		string 	   `json:"logo"`
}

type AplikasiInfo struct {
	ID			string     `json:"id"`
	Name		string	   `json:"name"`
	Logo		string 	   `json:"logo"`
}

type PermintaanInfo struct {
	ID                string                 `json:"id"`
	Pemda             PemdaInfo              `json:"pemda"`
	Aplikasi          AplikasiInfo           `json:"aplikasi"`
	Menu              string                 `json:"menu"`
	KondisiAwal       string                 `json:"kondisi_awal"`
	KondisiDiharapkan string                 `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time              `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time              `json:"tanggal_deadline"`
	Lampiran          permintaan.StringArray `json:"lampiran"`
}

type ProgrammerInfo struct {
	ID       	   string `json:"id"`
	Username 	   string `json:"username"`
	FullName 	   string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type LaporanInfo struct {
	ID              string        `json:"id"`
	LaporanProgress string        `json:"laporan_progress"`
	Status          string        `json:"status"`
	Programmer      ProgrammerInfo `json:"programmer"`
}

type VerifikatorInfo struct {
	ID       	   string `json:"id"`
	Username 	   string `json:"username"`
	FullName 	   string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type VerifikasiDetailResponse struct {
	ID             string          `json:"id"`
	Permintaan     PermintaanInfo  `json:"permintaan"`
	Laporan        LaporanInfo     `json:"laporan"`
	Verifikator    VerifikatorInfo `json:"verifikator"`
	Komentar       *string          `json:"komentar"`
	StatusVerified string          `json:"status_verified"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
