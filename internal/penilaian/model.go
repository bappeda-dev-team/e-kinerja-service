package penilaian

import "time"

type Penilaian struct {
	ID                  string    `json:"id"`
	DistribusiID        string    `json:"distribusi_id"`
	PenilaiID           string    `json:"penilai_id"`
	TingkatKeberhasilan int       `json:"tingkat_keberhasilan"`
	KetepatanWaktu      string    `json:"ketepatan_waktu"`
	Komentar            string    `json:"komentar"`
	TanggalSelesai      time.Time `json:"tanggal_selesai"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
