package laporan

import "time"

// === Request ===

type LaporanRequest struct {
	PermintaanID    string `json:"permintaan_id" validate:"required,uuid4"`
	LaporanProgress string `json:"laporan_progress" validate:"required,min=3"`
}

// === Response ===

type LaporanResponse struct {
	ID              string    `json:"id"`
	PermintaanID    string    `json:"permintaan_id"`
	ProgrammerID    string    `json:"programmer_id"`
	LaporanProgress string    `json:"laporan_progress"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
