package laporan

import (
	"aplikasi-internal/internal/permintaan"
	"time"
)

// === Request ===

type LaporanRequest struct {
	PermintaanID    string `json:"permintaan_id" validate:"required,uuid4"`
	LaporanProgress string `json:"laporan_progress" validate:"required,min=3"`
}

// === Response ===

type LaporanResponse struct {
	ID              string    `json:"id"`
	PermintaanID    string    `json:"permintaan_id"`
	ProgrammerID    string    `json:"programmer_id"`
	LaporanProgress string    `json:"laporan_progress"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
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

type LaporanDetailResponse struct {
	ID              string         `json:"id"`
	Permintaan      PermintaanInfo `json:"permintaan"`
	Programmer      ProgrammerInfo `json:"programmer"`
	LaporanProgress string         `json:"laporan_progress"`
	Status          string         `json:"status"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
