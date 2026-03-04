package laporan

import "time"

type Laporan struct {
	ID          	string    `json:"id"`
	PermintaanID    string    `json:"permintaan_id"`
	ProgrammerID 	string    `json:"programmer_id"`
	LaporanProgress string    `json:"laporan_progress"`
	CreatedAt   	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

type LaporanRequest struct {
	PermintaanID 	string `json:"permintaan_id" binding:"required"`
	ProgrammerID 	string `json:"programmer_id" binding:"required"`
	LaporanProgress string `json:"laporan_progress" binding:"required,min=3"`
}
