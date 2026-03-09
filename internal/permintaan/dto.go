package permintaan

import "time"

type PermintaanByNama struct {
	ID                string    `json:"id"`
	PemdaID           string    `json:"pemda"`
	AplikasiID        string    `json:"aplikasi"`
	Menu              string    `json:"menu"`
	KondisiAwal       string    `json:"kondisi_awal"`
	KondisiDiharapkan string    `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time `json:"tanggal_deadline"`
	CreatedBy         string    `json:"pembuat"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type PermintaanRequest struct {
	Menu              string    `json:"menu" validate:"required,min=3"`
	KondisiAwal       string    `json:"kondisi_awal" validate:"required,min=3"`
	KondisiDiharapkan string    `json:"kondisi_diharapkan" validate:"required,min=3"`
	TanggalPesanan    time.Time `json:"tanggal_pesanan" validate:"required"`
	TanggalDeadline   time.Time `json:"tanggal_deadline" validate:"required"`
}