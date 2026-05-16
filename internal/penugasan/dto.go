package penugasan

import "time"

// === Request ===

type CreatePenugasanRequest struct {
	DistribusiPelaksanaID string `json:"distribusi_pelaksana_id" validate:"required,uuid4"`
	Judul                 string `json:"judul" validate:"required,min=3"`
	Deskripsi             string `json:"deskripsi"`
	Deadline              string `json:"deadline"`
	Prioritas             string `json:"prioritas" validate:"required,oneof=low medium high"`
	EstimasiHari          *int   `json:"estimasi_hari"`
	Urutan                int    `json:"urutan"`
}

type UpdatePenugasanRequest struct {
	Judul        string `json:"judul" validate:"required,min=3"`
	Deskripsi    string `json:"deskripsi"`
	Deadline     string `json:"deadline"`
	Prioritas    string `json:"prioritas" validate:"required,oneof=low medium high"`
	EstimasiHari *int   `json:"estimasi_hari"`
	Urutan       int    `json:"urutan"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=belum_mulai sedang_berjalan selesai revisi"`
}

type ReassignRequest struct {
	DistribusiPelaksanaID string `json:"distribusi_pelaksana_id" validate:"required,uuid4"`
}

// === Response ===

type PelaksanaInfo struct {
	ID           string `json:"id"`
	ProgrammerID string `json:"programmer_id"`
	FullName     string `json:"full_name"`
}

type PenugasanResponse struct {
	ID                    string         `json:"id"`
	DistribusiPelaksanaID string         `json:"distribusi_pelaksana_id"`
	Pelaksana             *PelaksanaInfo `json:"pelaksana,omitempty"`
	Judul                 string         `json:"judul"`
	Deskripsi             string         `json:"deskripsi"`
	Deadline              *time.Time     `json:"deadline"`
	Prioritas             string         `json:"prioritas"`
	EstimasiHari          *int           `json:"estimasi_hari"`
	Urutan                int            `json:"urutan"`
	Status                string         `json:"status"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
}
