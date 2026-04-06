package all_activity

import "time"

type ActorInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type ActivityItem struct {
	ActivityType string    `json:"activity_type"` // permintaan | distribusi | pelaksana | laporan | verifikasi
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
