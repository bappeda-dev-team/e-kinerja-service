package routes

import (
	"aplikasi-internal/internal/distribusi"
	"aplikasi-internal/internal/laporan"
	"aplikasi-internal/internal/master_aplikasi"
	"aplikasi-internal/internal/master_pemda"
	"aplikasi-internal/internal/middleware"
	"aplikasi-internal/internal/pelaksana"
	"aplikasi-internal/internal/permintaan"
	"aplikasi-internal/internal/roles"
	"aplikasi-internal/internal/user"
	"aplikasi-internal/internal/verifikasi"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(middleware.ErrorHandler())

	r.GET("/roles", roles.GetRoles) // SELESAI
	
	r.POST("/logout", user.Logout)
	r.POST("/auth", user.Auth)
	r.POST("/register", user.Register)

	r.GET("/master-aplikasi", master_aplikasi.GetAplikasi) // SELESAI
	r.GET("/master-aplikasi/:id", master_aplikasi.GetAplikasiID) // SELESAI
	r.POST("/master-aplikasi", master_aplikasi.CreateAplikasi) // SELESAI
	r.PUT("/master-aplikasi/:id", master_aplikasi.UpdateAplikasi) // SELESAI
	r.DELETE("/master-aplikasi/:id", master_aplikasi.DeleteAplikasi) // SELESAI

	r.GET("/master-pemda", master_pemda.GetPemda) // SELESAI
	r.GET("/master-pemda/:id", master_pemda.GetPemdaID) // SELESAI
	r.POST("/master-pemda", master_pemda.CreatePemda) // SELESAI
	r.PUT("/master-pemda/:id", master_pemda.UpdatePemda) // SELESAI
	r.DELETE("/master-pemda/:id", master_pemda.DeletePemda) // SELESAI

	r.GET("/permintaan", permintaan.GetPermintaan) // SELESAI
	r.GET("/permintaan/:id", permintaan.GetPermintaanId) // SELESAI
	r.GET("/permintaan-nama", permintaan.GetPermintaanNama) // SELESAI
	r.GET("/permintaan-nama/:id", permintaan.GetPermintaanNamaId) // SELESAI
	r.POST("/permintaan", permintaan.CreatePermintaan) // SELESAI
	r.PUT("/permintaan/:id", permintaan.UpdatePermintaan) // SELESAI 
	r.DELETE("/permintaan/:id", permintaan.DeletePermintaan) // SELESAI

	r.GET("/distribusi", distribusi.GetDistribusi) // SELESAI
	r.GET("/distribusi/:id", distribusi.GetDistribusiId) // SELESAI
	r.POST("/distribusi", distribusi.CreateDistribusi) 
	r.PUT("/distribusi/:id", distribusi.UpdateDistribusi) 
	r.DELETE("/distribusi/:id", distribusi.DeleteDistribusi)  

	r.GET("/pelaksana", pelaksana.GetPelaksana) // SELESAI
	r.GET("/pelaksana/:id", pelaksana.GetPelaksanaID) // SELESAI
	r.POST("/pelaksana", pelaksana.CreatePelaksana) 
	r.PUT("/pelaksana/:id", pelaksana.UpdatePelaksana) 
	r.DELETE("/pelaksana/:id", pelaksana.DeletePelaksana) 

	r.GET("/laporan", laporan.GetLaporan) 
	r.GET("/laporan/:id", laporan.GetLaporanID) 
	r.POST("/laporan", laporan.CreateLaporan) 
	r.PUT("/laporan/:id", laporan.UpdateLaporan) 
	r.DELETE("/laporan/:id", laporan.DeleteLaporan) 

	r.GET("/verifikasi", verifikasi.GetVerifikasi) 
	r.POST("/verifikasi", verifikasi.CreateVerifikasi) 
}
