package superadmin_dashboard

import (
	"aplikasi-internal/internal/permintaan"
	"time"
)

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

type AdminInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type PelaksanaInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type DistribusiInfo struct {
	ID        string          `json:"id"`
	Admin     AdminInfo       `json:"admin"`
	Komentar  string          `json:"komentar"`
	Pelaksana []PelaksanaInfo `json:"pelaksana"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type LaporanInfo struct {
	ID              string      `json:"id"`
	Programmer      PembuatInfo `json:"programmer"`
	LaporanProgress string      `json:"laporan_progress"`
	Status          string      `json:"status"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type DashboardPermintaanItem struct {
	ID                string                 `json:"id"`
	Pemda             PemdaInfo              `json:"pemda"`
	Aplikasi          AplikasiInfo           `json:"aplikasi"`
	Menu              string                 `json:"menu"`
	KondisiAwal       string                 `json:"kondisi_awal"`
	KondisiDiharapkan string                 `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time              `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time              `json:"tanggal_deadline"`
	Lampiran          permintaan.StringArray `json:"lampiran"`
	Status            string                 `json:"status"`
	Pembuat           PembuatInfo            `json:"pembuat"`
	Distribusi        []DistribusiInfo       `json:"distribusi"`
	Laporan           []LaporanInfo          `json:"laporan"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
}

type DashboardResponse struct {
	TotalPermintaan int                       `json:"total_permintaan"`
	TotalDistribusi int                       `json:"total_distribusi"`
	TotalLaporan    int                       `json:"total_laporan"`
	Permintaan      []DashboardPermintaanItem `json:"permintaan"`
}
