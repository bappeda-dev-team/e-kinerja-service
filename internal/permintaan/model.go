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

type PermintaanByNama struct {
	ID          	  string    `json:"id"`
	PemdaID        	  string    `json:"pemda"`
	AplikasiID 		  string    `json:"aplikasi"`
	Menu 			  string    `json:"menu"`
	KondisiAwal 	  string    `json:"kondisi_awal"`
	KondisiDiharapkan string    `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time `json:"tanggal_deadline"`
	CreatedBy 		  string    `json:"pembuat"`
	CreatedAt   	  time.Time `json:"created_at"`
	UpdatedAt   	  time.Time `json:"updated_at"`
}

type PermintaanRequest struct {
	PemdaID 		  string 	`json:"pemda_id" binding:"required"`
	AplikasiID 		  string 	`json:"aplikasi_id" binding:"required"`
	Menu 			  string 	`json:"menu" binding:"required,min=3"`
	KondisiAwal 	  string 	`json:"kondisi_awal" binding:"required,min=3"`
	KondisiDiharapkan string 	`json:"kondisi_diharapkan" binding:"required,min=3"`
	TanggalPesanan 	  time.Time `json:"tanggal_pesanan" binding:"required"`
	TanggalDeadline   time.Time `json:"tanggal_deadline" binding:"required"`
	CreatedBy 		  string 	`json:"created_by" binding:"required"`
}