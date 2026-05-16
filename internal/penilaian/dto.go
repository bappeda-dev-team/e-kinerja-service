package penilaian

import "time"

type CreatePenilaianRequest struct {
	DistribusiID        string `json:"distribusi_id" validate:"required,uuid4"`
	TingkatKeberhasilan int    `json:"tingkat_keberhasilan" validate:"required,min=0,max=100"`
	Komentar            string `json:"komentar"`
	TanggalSelesai      string `json:"tanggal_selesai" validate:"required"`
}

type UpdatePenilaianRequest struct {
	TingkatKeberhasilan *int   `json:"tingkat_keberhasilan" validate:"omitempty,min=0,max=100"`
	Komentar            string `json:"komentar"`
	TanggalSelesai      string `json:"tanggal_selesai"`
}

type PenilaiInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type DistribusiInfo struct {
	ID           string `json:"id"`
	PermintaanID string `json:"permintaan_id"`
	Menu         string `json:"menu"`
}

type PenilaianResponse struct {
	ID                  string         `json:"id"`
	Distribusi          DistribusiInfo `json:"distribusi"`
	Penilai             PenilaiInfo    `json:"penilai"`
	TingkatKeberhasilan int            `json:"tingkat_keberhasilan"`
	KetepatanWaktu      string         `json:"ketepatan_waktu"`
	Komentar            string         `json:"komentar"`
	TanggalSelesai      time.Time      `json:"tanggal_selesai"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
}
