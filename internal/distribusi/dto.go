package distribusi

import (
	"aplikasi-internal/internal/permintaan"
	"time"
)

// === Request ===

type DistribusiRequest struct {
	PermintaanID string `json:"permintaan_id" validate:"required,uuid4"`
	Komentar     string `json:"komentar"`
}

// === Response ===

type DistribusiResponse struct {
	ID           string    `json:"id"`
	PermintaanID string    `json:"permintaan_id"`
	AdminID      string    `json:"admin_id"`
	Komentar     string    `json:"komentar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PermintaanInfo struct {
	ID                string                 `json:"id"`
	Pemda             string                 `json:"pemda"`
	Aplikasi          string                 `json:"aplikasi"`
	Menu              string                 `json:"menu"`
	KondisiAwal       string                 `json:"kondisi_awal"`
	KondisiDiharapkan string                 `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time              `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time              `json:"tanggal_deadline"`
	Lampiran          permintaan.StringArray `json:"lampiran"`
}

type AdminInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type PelaksanaInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type DistribusiDetailResponse struct {
	ID         string          `json:"id"`
	Permintaan PermintaanInfo  `json:"permintaan"`
	Admin      AdminInfo       `json:"admin"`
	Komentar   string          `json:"komentar"`
	Pelaksana  []PelaksanaInfo `json:"pelaksana"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
