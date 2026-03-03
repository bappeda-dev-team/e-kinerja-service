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

	r.GET("/roles", roles.GetRoles) // 													    		SELESAI
	
	r.GET("/register", user.Register) //															SELESAI
	r.GET("/logout", user.Logout)
	r.POST("/auth", user.Auth)
	r.POST("/register", user.CreateRegister)

	r.GET("/master-aplikasi", master_aplikasi.GetAplikasi) // hanya superadmin 						SELESAI
	r.GET("/master-aplikasi/:id", master_aplikasi.GetAplikasiID) // hanya superadmin				SELESAI
	r.POST("/master-aplikasi", master_aplikasi.CreateAplikasi) // hanya superadmin					SELESAI
	r.PUT("/master-aplikasi/:id", master_aplikasi.UpdateAplikasi) // hanya superadmin				SELESAI
	r.DELETE("/master-aplikasi/:id", master_aplikasi.DeleteAplikasi) // hanya superadmin			SELESAI

	r.GET("/master-pemda", master_pemda.GetPemda) // hanya superadmin								SELESAI
	r.GET("/master-pemda/:id", master_pemda.GetPemdaID) // hanya superadmin							SELESAI
	r.POST("/master-pemda", master_pemda.CreatePemda) // hanya superadmin							SELESAI
	r.PUT("/master-pemda/:id", master_pemda.UpdatePemda) // hanya superadmin						SELESAI
	r.DELETE("/master-pemda/:id", master_pemda.DeletePemda) // hanya superadmin						SELESAI

	r.GET("/permintaan", permintaan.GetPermintaan) // semua users dapat mengakses					SELESAI
	r.GET("/permintaan/:id", permintaan.GetPermintaanId) // semua users dapat mengakses				SELESAI
	r.POST("/permintaan", permintaan.CreatePermintaan) // hanya superadmin
	r.PUT("/permintaan/:id", permintaan.UpdatePermintaan) // hanya superadmin
	r.DELETE("/permintaan/:id", permintaan.DeletePermintaan) // hanya superadmin

	r.GET("/distribusi", distribusi.GetDistribusi) // semua users dapat mengakses					SELESAI
	r.GET("/distribusi/:id", distribusi.GetDistribusiId) // semua users dapat mengakses				SELESAI
	r.POST("/distribusi", distribusi.CreateDistribusi) // hanya superadmin 
	r.PUT("/distribusi/:id", distribusi.UpdateDistribusi) // hanya superadmin 
	r.DELETE("/distribusi/:id", distribusi.DeleteDistribusi) // hanya superadmin 

	r.GET("/pelaksana", pelaksana.GetPelaksana) // semua users dapat mengakses						SELESAI
	r.GET("/pelaksana/:id", pelaksana.GetPelaksanaID) // semua users dapat mengakses				SELESAI
	r.POST("/pelaksana", pelaksana.CreatePelaksana) // hanya superadmin & admin 
	r.PUT("/pelaksana/:id", pelaksana.UpdatePelaksana) // hanya superadmin & admin 
	r.DELETE("/pelaksana/:id", pelaksana.DeletePelaksana) // hanya superadmin & admin

	r.GET("/laporan", laporan.GetLaporan) // semua users dapat mengakses
	r.GET("/laporan/:id", laporan.GetLaporanID) // semua users dapat mengakses
	r.POST("/laporan", laporan.CreateLaporan) // hanya superadmin, programmer & level 2
	r.PUT("/laporan/:id", laporan.UpdateLaporan) // hanya superadmin, programmer & level 2
	r.DELETE("/laporan/:id", laporan.DeleteLaporan) // hanya superadmin, programmer & level 2

	r.GET("/verifikasi", verifikasi.GetVerifikasi) // semua users dapat mengakses
	r.POST("/verifikasi", verifikasi.CreateVerifikasi) // hanya level 2
}
