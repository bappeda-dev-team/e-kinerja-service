package master_aplikasi

import "time"

type MasterAplikasi struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateMasterAplikasiRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type ConfirmLogoRequest struct {
	Key string `json:"key" validate:"required"`
}