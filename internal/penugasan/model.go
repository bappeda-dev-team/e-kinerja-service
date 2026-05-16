package penugasan

import "time"

type Penugasan struct {
	ID                    string     `json:"id"`
	DistribusiPelaksanaID string     `json:"distribusi_pelaksana_id"`
	Judul                 string     `json:"judul"`
	Deskripsi             string     `json:"deskripsi"`
	Deadline              *time.Time `json:"deadline"`
	Prioritas             string     `json:"prioritas"`
	EstimasiHari          *int       `json:"estimasi_hari"`
	Urutan                int        `json:"urutan"`
	Status                string     `json:"status"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}
