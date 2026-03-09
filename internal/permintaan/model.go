package permintaan

import "time"

type Permintaan struct {
	ID          	  string    `json:"id"`
	PemdaID        	  string    `json:"pemda_id"`
	AplikasiID 		  string    `json:"aplikasi_id"`
	Menu 			  string    `json:"menu"`
	KondisiAwal 	  string    `json:"kondisi_awal"`
	KondisiDiharapkan string    `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time `json:"tanggal_deadline"`
	CreatedBy 		  string    `json:"created_by"`
	CreatedAt   	  time.Time `json:"created_at"`
	UpdatedAt   	  time.Time `json:"updated_at"`
}

