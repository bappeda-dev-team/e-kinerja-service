package distribusi

import "time"

// entitas DB
type Distribusi struct {
	ID           string    `json:"id"`
	PermintaanID string    `json:"permintaan_id"`
	AdminID      string    `json:"admin_id"`
	Komentar     string    `json:"komentar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type KomentarDistribusi struct {
	ID           string    `json:"id"`
	DistribusiID string    `json:"distribusi_id"`
	UserID 		 string    `json:"user_id"`
	Komentars    string    `json:"komentars"`
	IsRead       bool      `json:"is_read"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
