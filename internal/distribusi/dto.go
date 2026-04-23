package distribusi

import (
	"aplikasi-internal/internal/permintaan"
	"errors"
	"time"

	"github.com/google/uuid"
)

// === Request ===

type DistribusiRequest struct {
	PermintaanID  string   `json:"permintaan_id" validate:"required,uuid4"`
	Komentar      string   `json:"komentar"`
	ProgrammerIDs []string `json:"programmer_ids"`
	Pelaksana     []string `json:"pelaksana"`
}

type KomentarDistribusiRequest struct {
	Komentars      string   `json:"komentars"`
}

func (r *DistribusiRequest) NormalizePelaksana() {
	ids := r.ProgrammerIDs
	if len(r.Pelaksana) > 0 {
		ids = r.Pelaksana
	}

	if len(ids) == 0 {
		return
	}

	r.ProgrammerIDs = append([]string(nil), ids...)
	r.Pelaksana = append([]string(nil), ids...)
}

func (r DistribusiRequest) ValidatePelaksana() error {
	if len(r.ProgrammerIDs) == 0 {
		return errors.New("pelaksana atau programmer_ids harus diisi")
	}

	for _, programmerID := range r.ProgrammerIDs {
		if _, err := uuid.Parse(programmerID); err != nil {
			return errors.New("pelaksana atau programmer_ids harus berupa daftar UUID yang valid")
		}
	}

	return nil
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

type PermintaanInfo struct {
	ID                string                 `json:"id"`
	Pemda             PemdaInfo              `json:"pemda"`
	Aplikasi          AplikasiInfo           `json:"aplikasi"`
	Menu              string                 `json:"menu"`
	KondisiAwal       string                 `json:"kondisi_awal"`
	KondisiDiharapkan string                 `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time              `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time              `json:"tanggal_deadline"`
	Lampiran          permintaan.StringArray `json:"lampiran"`
}

type AdminInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type PelaksanaInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type VerifikasiInfo struct {
	ID             string `json:"id"`
	Komentar       string `json:"komentar"`
	StatusVerified string `json:"status_verified"`
}

type KomentarInfo struct {
	ID         string          `json:"id"`
	FullName   string          `json:"full_name"`
	Komentars  string          `json:"komentar"`
	CreatedAt  time.Time       `json:"created_at"`
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

type DistribusiFullResponse struct {
	ID         string          `json:"id"`
	Permintaan PermintaanInfo  `json:"permintaan"`
	Admin      AdminInfo       `json:"admin"`
	Komentar   []KomentarInfo  `json:"komentars"`
	Pelaksana  []PelaksanaInfo `json:"pelaksana"`
	Verifikasi VerifikasiInfo  `json:"verifikasi"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

type KomentarResponse struct {
	ID         string          `json:"id"`
	FullName   string          `json:"full_name"`
	Komentars  string          `json:"komentars"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
