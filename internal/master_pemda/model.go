package master_pemda

import "time"

type MasterPemda struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MasterPemdaRequest struct {
	Name string `json:"name" form:"name" validate:"required,min=3"`
}
