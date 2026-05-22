package pelaksana

import "time"

// === Request ===

type PelaksanaRequest struct {
	DistribusiID string `json:"distribusi_id" validate:"required,uuid4"`
	ProgrammerID string `json:"programmer_id" validate:"required,uuid4"`
}

// === Response ===

type PelaksanaResponse struct {
	ID           string    `json:"id"`
	DistribusiID string    `json:"distribusi_id"`
	ProgrammerID string    `json:"programmer_id"`
	IsRead       bool      `json:"is_read"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type KomentarInfo struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Komentar  string    `json:"komentar"`
	CreatedAt time.Time `json:"created_at"`
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

type PermintaanLite struct {
	ID              string    `json:"id"`
	TanggalDeadline time.Time `json:"tanggal_deadline"`
}

type DistribusiInfo struct {
	ID           string         `json:"id"`
	PermintaanID string         `json:"permintaan_id"`
	Pemda        PemdaInfo      `json:"pemda"`
	Aplikasi     AplikasiInfo   `json:"aplikasi"`
	Permintaan   PermintaanLite `json:"permintaan"`
	Komentars    []KomentarInfo `json:"komentars"`
}

type ProgrammerInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type PelaksanaDetailResponse struct {
	ID         string         `json:"id"`
	Distribusi DistribusiInfo `json:"distribusi"`
	Programmer ProgrammerInfo `json:"programmer"`
	IsRead     bool           `json:"is_read"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}
