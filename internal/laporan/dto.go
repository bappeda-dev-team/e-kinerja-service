package laporan

import (
	"aplikasi-internal/internal/permintaan"
	"time"
)

// === Request ===

type LaporanRequest struct {
	PermintaanID    string `json:"permintaan_id" form:"permintaan_id" validate:"required,uuid4"`
	PenugasanID     string `json:"penugasan_id" form:"penugasan_id"`
	LaporanProgress string `json:"laporan_progress" form:"laporan_progress" validate:"required,min=3"`
	Status          string `json:"status" form:"status" validate:"required"`
}

type LaporanUpdateRequest struct {
	PermintaanID    	  string `json:"permintaan_id" validate:"required,uuid4"`
	LaporanProgress 	  string `json:"laporan_progress" validate:"required,min=3"`
	Status          	  string `json:"status" validate:"required"`
	VerifikasiID    	  string `json:"verifikasi_id"`
	StatusVerified  	  string `json:"status_verified"`
	IsSubmittedToVerified bool   `json:"is_submitted_to_verified"`
}


type VerifikasiRequest struct {
	LaporanID string `json:"laporan_id" validate:"required,uuid4"`
}

type KomentarLaporanRequest struct {
	Komentar      string   `json:"komentar"`
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
	ID   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type AplikasiInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
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
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type VerifikasiInfo struct {
	ID             		  string    `json:"id"`
	StatusVerified 		  string    `json:"status_verified"`
	IsSubmittedToVerified bool   	`json:"is_submitted_to_verified"`
	CreatedAt      		  time.Time `json:"created_at"`
	UpdatedAt      		  time.Time `json:"updated_at"`
}

type KomentarInfo struct {
	ID         string          `json:"id"`
	FullName   string          `json:"full_name"`
	Komentar   string          `json:"komentar"`
	CreatedAt  time.Time       `json:"created_at"`
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

type LaporanVerifikasiResponse struct {
	ID              string         `json:"id"`
	Permintaan      PermintaanInfo `json:"permintaan"`
	LaporanProgress string         `json:"laporan_progress"`
	Status          string         `json:"status"`
}

type VerifikasiResponse struct {
	ID             string                    `json:"id"`
	Laporan        LaporanVerifikasiResponse `json:"laporan"`
	Programmer     ProgrammerInfo            `json:"programmer"`
	StatusVerified string                    `json:"status_verified"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`
}

type LaporanFullResponse struct {
	ID              string           `json:"id"`
	Permintaan      PermintaanInfo   `json:"permintaan"`
	Programmer      ProgrammerInfo   `json:"programmer"`
	PenugasanID     *string          `json:"penugasan_id"`
	LaporanProgress string           `json:"laporan_progress"`
	Status          string           `json:"status"`
	Lampiran        StringArray      `json:"lampiran"`
	Verifikasi      []VerifikasiInfo `json:"verifikasi"`
	Komentars       []KomentarInfo   `json:"komentars"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type HistoryResponse struct {
	ID 			       string	 `json:"id"`
	LaporanID 	   	   string	 `json:"laporan_id"`
	ProgrammerID 	   string	 `json:"programmer_id"`
	OldStatus	 	   string	 `json:"old_status"`
	NewStatus	 	   string	 `json:"new_status"`
	OldProgress	 	   string	 `json:"old_progress"`
	NewProgress	 	   string	 `json:"new_progress"`
	CreatedAt          time.Time `json:"created_at"`
}

type KomentarResponse struct {
	ID         string          `json:"id"`
	FullName   string          `json:"full_name"`
	Komentar   string          `json:"komentar"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
