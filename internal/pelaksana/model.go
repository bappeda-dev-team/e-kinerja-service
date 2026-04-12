package pelaksana

import "time"

// entitas DB
type Pelaksana struct {
	ID           string    `json:"id"`
	DistribusiID string    `json:"distribusi_id"`
	ProgrammerID string    `json:"programmer_id"`
	IsRead       bool      `json:"is_read"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
