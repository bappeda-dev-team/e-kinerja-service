package permintaan

import (
	"time"
)

// === Request ===

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

type PermintaanRequest struct {
	PemdaID           string `json:"pemda_id"           form:"pemda_id"           validate:"required,uuid4"`
	AplikasiID        string `json:"aplikasi_id"        form:"aplikasi_id"        validate:"required,uuid4"`
	Menu              string `json:"menu"               form:"menu"               validate:"required,min=3"`
	KondisiAwal       string `json:"kondisi_awal"       form:"kondisi_awal"       validate:"required,min=3"`
	KondisiDiharapkan string `json:"kondisi_diharapkan" form:"kondisi_diharapkan" validate:"required,min=3"`
	TanggalPesanan    string `json:"tanggal_pesanan"    form:"tanggal_pesanan"    validate:"required"`
	TanggalDeadline   string `json:"tanggal_deadline"   form:"tanggal_deadline"   validate:"required"`
}

// === Response ===

type PermintaanResponse struct {
	ID                string      `json:"id"`
	PemdaID           string      `json:"pemda_id"`
	AplikasiID        string      `json:"aplikasi_id"`
	Menu              string      `json:"menu"`
	KondisiAwal       string      `json:"kondisi_awal"`
	KondisiDiharapkan string      `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time   `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time   `json:"tanggal_deadline"`
	Lampiran          StringArray `json:"lampiran"`
	Status            string      `json:"status"`
	CreatedBy         string      `json:"created_by"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
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

type PembuatInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type PermintaanDetailResponse struct {
	ID                string       `json:"id"`
	Pemda             PemdaInfo    `json:"pemda"`
	Aplikasi          AplikasiInfo `json:"aplikasi"`
	Menu              string       `json:"menu"`
	KondisiAwal       string       `json:"kondisi_awal"`
	KondisiDiharapkan string       `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time    `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time    `json:"tanggal_deadline"`
	Lampiran          StringArray  `json:"lampiran"`
	Status            string       `json:"status"`
	Pembuat           PembuatInfo  `json:"pembuat"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
}
