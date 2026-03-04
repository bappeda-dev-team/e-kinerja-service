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
	r.GET("/roles/:id", roles.GetRoleID) // SELESAI
	
	r.POST("/logout", user.Logout)
	r.POST("/auth", user.Auth)
	r.GET("/user", user.GetAllUser) // SELESAI
	r.GET("/user/:id", user.GetUserID) // SELESAI
	r.POST("/create-user", user.Create) //SELESAI

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
	r.GET("/distribusi/:id", distribusi.GetDistribusiById) // SELESAI
	r.GET("/distribusi-nama", distribusi.GetDistribusiByNama) // SELESAI
	r.GET("/distribusi-nama/:id", distribusi.GetDistribusiByNamaId) // SELESAI
	r.POST("/distribusi", distribusi.CreateDistribusi) // SELESAI
	r.PUT("/distribusi/:id", distribusi.UpdateDistribusi) // SELESAI
	r.DELETE("/distribusi/:id", distribusi.DeleteDistribusi) // SELESAI 

	r.GET("/pelaksana", pelaksana.GetPelaksana) // SELESAI
	r.GET("/pelaksana/:id", pelaksana.GetPelaksanaID) // SELESAI
	r.GET("/pelaksana-nama", pelaksana.GetPelaksanaByNama) // SELESAI
	r.GET("/pelaksana-nama/:id", pelaksana.GetPelaksanaByNamaID) // SELESAI
	r.POST("/pelaksana", pelaksana.CreatePelaksana) // SELESAI
	r.PUT("/pelaksana/:id", pelaksana.UpdatePelaksana) // SELESAI
	r.DELETE("/pelaksana/:id", pelaksana.DeletePelaksana) // SELESAI

	r.GET("/laporan", laporan.GetLaporan) // SELESAI
	r.GET("/laporan/:id", laporan.GetLaporanID) // SELESAI
	r.POST("/laporan", laporan.CreateLaporan) // SELESAI
	r.PUT("/laporan/:id", laporan.UpdateLaporan) // SELESAI
	r.DELETE("/laporan/:id", laporan.DeleteLaporan) // SELESAI

	r.GET("/verifikasi", verifikasi.GetVerifikasi) // SELESAI
	r.GET("/verifikasi/:id", verifikasi.GetVerifikasiID) // SELESAI
	r.POST("/verifikasi", verifikasi.CreateVerifikasi) // SELESAI
	r.PUT("/verifikasi/:id", verifikasi.UpdateVerifikasi) // SELESAI
	r.DELETE("/verifikasi/:id", verifikasi.DeleteVerifikasi) // SELESAI
}
