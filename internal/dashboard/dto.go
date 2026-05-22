package dashboard

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// StringArray adalah local copy dari permintaan.StringArray untuk scan kolom TEXT[] PostgreSQL.
type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	if s == nil {
		return "{}", nil
	}
	result := "{"
	for i, v := range s {
		if i > 0 {
			result += ","
		}
		b, _ := json.Marshal(v)
		result += string(b)
	}
	result += "}"
	return result, nil
}

func (s *StringArray) Scan(src interface{}) error {
	if src == nil {
		*s = StringArray{}
		return nil
	}
	var str string
	switch v := src.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return fmt.Errorf("StringArray.Scan: unsupported type %T", src)
	}
	if str == "{}" || str == "" {
		*s = StringArray{}
		return nil
	}
	str = str[1 : len(str)-1]
	if str == "" {
		*s = StringArray{}
		return nil
	}
	var result []string
	if err := json.Unmarshal([]byte("["+str+"]"), &result); err != nil {
		*s = splitPGArray(str)
		return nil
	}
	*s = result
	return nil
}

func splitPGArray(s string) []string {
	var result []string
	var cur string
	inQuote := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '"' {
			inQuote = !inQuote
		} else if c == ',' && !inQuote {
			result = append(result, cur)
			cur = ""
		} else {
			cur += string(c)
		}
	}
	if cur != "" {
		result = append(result, cur)
	}
	return result
}

// ── Shared sub-types ──────────────────────────────────────────────────────────

type UserInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
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

type DistribusiItem struct {
	ID        string     `json:"id"`
	Admin     UserInfo   `json:"admin"`
	Komentar  string     `json:"komentar"`
	Pelaksana []UserInfo `json:"pelaksana"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type LaporanItem struct {
	ID              string    `json:"id"`
	Programmer      UserInfo  `json:"programmer"`
	LaporanProgress string    `json:"laporan_progress"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type PermintaanItem struct {
	ID                string           `json:"id"`
	Pemda             PemdaInfo        `json:"pemda"`
	Aplikasi          AplikasiInfo     `json:"aplikasi"`
	Menu              string           `json:"menu"`
	KondisiAwal       string           `json:"kondisi_awal"`
	KondisiDiharapkan string           `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time        `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time        `json:"tanggal_deadline"`
	Lampiran          StringArray      `json:"lampiran"`
	Status            string           `json:"status"`
	Pembuat           UserInfo         `json:"pembuat"`
	Distribusi        []DistribusiItem `json:"distribusi"`
	Laporan           []LaporanItem    `json:"laporan"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}

// ── Per-role response types ───────────────────────────────────────────────────

type SuperAdminDashboardResponse struct {
	TotalPermintaan int              `json:"total_permintaan"`
	TotalDistribusi int              `json:"total_distribusi"`
	TotalLaporan    int              `json:"total_laporan"`
	Permintaan      []PermintaanItem `json:"permintaan"`
}

type AdminDashboardResponse struct {
	TotalPermintaan int              `json:"total_permintaan"`
	TotalDistribusi int              `json:"total_distribusi"`
	TotalPelaksana  int              `json:"total_pelaksana"`
	Permintaan      []PermintaanItem `json:"permintaan"`
}

type PenugasanItem struct {
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

type ProgrammerDashboardResponse struct {
	TotalPenugasan int             `json:"total_penugasan"`
	TotalLaporan   int             `json:"total_laporan"`
	Penugasan      []PenugasanItem `json:"penugasan"`
	Laporan        []LaporanItem   `json:"laporan"`
}

type VerifikatorLaporanItem struct {
	ID              string       `json:"id"`
	Permintaan      PermintaanItem `json:"permintaan"`
	Programmer      UserInfo     `json:"programmer"`
	LaporanProgress string       `json:"laporan_progress"`
	Status          string       `json:"status"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

type VerifikatorDashboardResponse struct {
	TotalMenunggu int                      `json:"total_menunggu"`
	Laporan       []VerifikatorLaporanItem `json:"laporan"`
}

// ── Activity feed types ───────────────────────────────────────────────────────

type ActorInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type ActivityItem struct {
	ActivityType string    `json:"activity_type"` // permintaan|distribusi|pelaksana|penugasan|laporan|verifikasi|komentar_distribusi|komentar_laporan|penilaian
	ActivityID   string    `json:"activity_id"`
	RefID        string    `json:"ref_id"` // permintaan_id sebagai titik referensi
	Actor        ActorInfo `json:"actor"`
	PemdaName    string    `json:"pemda_name"`
	AplikasiName string    `json:"aplikasi_name"`
	Menu         string    `json:"menu"`
	Status       string    `json:"status"`
	Komentar     string    `json:"komentar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
