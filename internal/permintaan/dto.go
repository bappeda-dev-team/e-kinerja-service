package permintaan

import "time"

// === Request ===

type PermintaanRequest struct {
	PemdaID           string    `json:"pemda_id" validate:"required,uuid4"`
	AplikasiID        string    `json:"aplikasi_id" validate:"required,uuid4"`
	Menu              string    `json:"menu" validate:"required,min=3"`
	KondisiAwal       string    `json:"kondisi_awal" validate:"required,min=3"`
	KondisiDiharapkan string    `json:"kondisi_diharapkan" validate:"required,min=3"`
	TanggalPesanan    time.Time `json:"tanggal_pesanan" validate:"required"`
	TanggalDeadline   time.Time `json:"tanggal_deadline" validate:"required"`
}

type ConfirmLampiranRequest struct {
	// Keys adalah daftar object key S3 hasil upload (max 3)
	Keys []string `json:"keys" validate:"required,min=1,max=3"`
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
	CreatedBy         string      `json:"created_by"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

type PermintaanDetailResponse struct {
	ID                string      `json:"id"`
	Pemda             string      `json:"pemda"`
	Aplikasi          string      `json:"aplikasi"`
	Menu              string      `json:"menu"`
	KondisiAwal       string      `json:"kondisi_awal"`
	KondisiDiharapkan string      `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time   `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time   `json:"tanggal_deadline"`
	Lampiran          StringArray `json:"lampiran"`
	Pembuat           string      `json:"pembuat"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}