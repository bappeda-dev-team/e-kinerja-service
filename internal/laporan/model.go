package laporan

import "time"

type Laporan struct {
	ID              string    `json:"id"`
	PermintaanID    string    `json:"permintaan_id"`
	ProgrammerID    string    `json:"programmer_id"`
	LaporanProgress string    `json:"laporan_progress"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Verifikasi struct {
	ID             string    `json:"id"`
	LaporanID      string    `json:"laporan_id"`
	ProgrammerID   string    `json:"programmer_id"`
	Komentar       string    `json:"komentar"`
	StatusVerified string    `json:"status_verified"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type History struct {
	ID 			       string	 `json:"id"`
	LaporanID 	   	   string	 `json:"laporan_id"`
	ProgrammerID 	   string	 `json:"programmer_id"`
	OldStatus	 	   string	 `json:"old_status"`
	NewStatus	 	   string	 `json:"new_status"`
	OldProgress	 	   string	 `json:"old_progress"`
	NewProgress	 	   string	 `json:"new_progress"`
	CreatedAt          time.Time `json:"created_at"`
}
