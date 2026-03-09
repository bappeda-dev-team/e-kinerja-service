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

	ma := r.Group("/master-aplikasi")
	ma.Use(middle_ware.JWTMiddleware)

	ma.GET("", master_aplikasi.GetAplikasi) // SELESAI
	ma.GET("/:id", master_aplikasi.GetAplikasiID) // SELESAI
	ma.POST("", master_aplikasi.CreateAplikasi) // SELESAI
	ma.PUT("/:id", master_aplikasi.UpdateAplikasi) // SELESAI
	ma.DELETE("/:id", master_aplikasi.DeleteAplikasi) // SELESAI

	mp := r.Group("/master-pemda")
	mp.Use(middle_ware.JWTMiddleware)

	mp.GET("", master_pemda.GetPemda) // SELESAI
	mp.GET("/:id", master_pemda.GetPemdaID) // SELESAI
	mp.POST("", master_pemda.CreatePemda) // SELESAI
	mp.PUT("/:id", master_pemda.UpdatePemda) // SELESAI
	mp.DELETE("/:id", master_pemda.DeletePemda) // SELESAI

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

	dp := r.Group("/pelaksana")
	dp.Use(middle_ware.JWTMiddleware)

	dp.GET("", pelaksana.GetPelaksana) // SELESAI
	dp.GET("/:id", pelaksana.GetPelaksanaID) // SELESAI
	dp.GET("/nama", pelaksana.GetPelaksanaByNama) // SELESAI
	dp.GET("/nama/:id", pelaksana.GetPelaksanaByNamaID) // SELESAI
	dp.POST("/distribusi/:distribusi_id/programmer/:programmer_id", pelaksana.CreatePelaksana) // SELESAI
	dp.PUT("/distribusi/:distribusi_id/programmer/:programmer_id/id/:id", pelaksana.UpdatePelaksana) // SELESAI
	dp.DELETE("/:id", pelaksana.DeletePelaksana) // SELESAI

	l := r.Group("/laporan")
	l.Use(middle_ware.JWTMiddleware)

	l.GET("", laporan.GetLaporan) // SELESAI
	l.GET("/:id", laporan.GetLaporanID) // SELESAI
	l.POST("/permintaan/:permintaan_id", laporan.CreateLaporan) // SELESAI
	l.PUT("/permintaan/:permintaan_id/id/:id", laporan.UpdateLaporan) // SELESAI
	l.DELETE("/:id", laporan.DeleteLaporan) // SELESAI

	v := r.Group("/verifikasi")
	v.Use(middle_ware.JWTMiddleware)

	v.GET("", verifikasi.GetVerifikasi) // SELESAI
	v.GET("/:id", verifikasi.GetVerifikasiID) // SELESAI
	v.POST("/laporan/:laporan_id", verifikasi.CreateVerifikasi) // SELESAI
	v.PUT("/laporan/:laporan_id/id/:id", verifikasi.UpdateVerifikasi) // SELESAI
	v.DELETE("/:id", verifikasi.DeleteVerifikasi) // SELESAI
}
