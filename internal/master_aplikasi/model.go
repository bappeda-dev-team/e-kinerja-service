package master_aplikasi

import "time"

type MasterAplikasi struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Logo      string    `json:"logo"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateMasterAplikasiRequest struct {
	Name string `json:"name" form:"name" validate:"required,min=3"`
	Link string `json:"link" form:"link"`
}
