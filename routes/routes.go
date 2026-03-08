package routes

import (
	"aplikasi-internal/internal/distribusi"
	"aplikasi-internal/internal/laporan"
	"aplikasi-internal/internal/master_aplikasi"
	"aplikasi-internal/internal/master_pemda"
	"aplikasi-internal/internal/middle_ware"
	"aplikasi-internal/internal/pelaksana"
	"aplikasi-internal/internal/permintaan"
	"aplikasi-internal/internal/roles"
	"aplikasi-internal/internal/user"
	"aplikasi-internal/internal/verifikasi"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "aplikasi-internal/docs"
)

func SetupRoutes(r *echo.Echo) {

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	// global error handler
	r.HTTPErrorHandler = middle_ware.ErrorHandler

	r.GET("/roles", roles.GetRoles) // SELESAI
	r.GET("/roles/:id", roles.GetRoleID) // SELESAI
	
	// e.POST("/logout", user.Logout)
	auth := r.Group("/auth")
	auth.POST("/login", user.Login)

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

	p := r.Group("/permintaan")
	p.Use(middle_ware.JWTMiddleware)

	p.GET("", permintaan.GetPermintaan) // SELESAI
	p.GET("/:id", permintaan.GetPermintaanId) // SELESAI
	p.GET("/nama", permintaan.GetPermintaanNama) // SELESAI
	p.GET("/nama/:id", permintaan.GetPermintaanNamaId) // SELESAI
	p.POST("/pemda/:pemda_id/aplikasi/:aplikasi_id", permintaan.CreatePermintaan) // SELESAI
	p.PUT("/pemda/:pemda_id/aplikasi/:aplikasi_id/id/:id", permintaan.UpdatePermintaan) // SELESAI 
	p.DELETE("/:id", permintaan.DeletePermintaan) // SELESAI

	d := r.Group("/distribusi")
	d.Use(middle_ware.JWTMiddleware)

	d.GET("", distribusi.GetDistribusi) // SELESAI
	d.GET("/:id", distribusi.GetDistribusiById) // SELESAI
	d.GET("/nama", distribusi.GetDistribusiByNama) // SELESAI
	d.GET("/nama/:id", distribusi.GetDistribusiByNamaId) // SELESAI
	d.POST("/permintaan/:permintaan_id", distribusi.CreateDistribusi) // SELESAI
	d.PUT("/permintaan/:permintaan_id/id/:id", distribusi.UpdateDistribusi) // SELESAI
	d.DELETE("/:id", distribusi.DeleteDistribusi) // SELESAI 

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
