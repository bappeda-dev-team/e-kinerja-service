package routes

import (
	"aplikasi-internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/login", controllers.Login)
	r.GET("/register", controllers.Register)
	r.GET("/logout", controllers.Logout)
	r.POST("/auth", controllers.Auth)
	r.POST("/register", controllers.CreateRegister)

	r.GET("/permintaan", controllers.GetPermintaan) // semua users dapat mengakses
	r.GET("/permintaan/:id", controllers.GetPermintaanId) // semua users dapat mengakses
	r.POST("/permintaan", controllers.CreatePermintaan) // hanya superadmin
	r.PUT("/permintaan/:id", controllers.UpdatePermintaan) // hanya superadmin
	r.DELETE("/permintaan/:id", controllers.DeletePermintaan) // hanya superadmin

	r.GET("/distribusi", controllers.GetDistribusi) // semua users dapat mengakses
	r.GET("/distribusi/:id", controllers.GetDistribusiId) // semua users dapat mengakses
	r.POST("/distribusi", controllers.CreateDistribusi) // hanya superadmin & admin
	r.PUT("/distribusi/:id", controllers.UpdateDistribusi) // hanya superadmin & admin
	r.DELETE("/distribusi/:id", controllers.DeleteDistribusi) // hanya superadmin & admin

	r.GET("/laporan", controllers.GetLaporan) // semua users dapat mengakses
	r.GET("/laporan/:id", controllers.GetLaporanID) // semua users dapat mengakses
	r.POST("/laporan", controllers.CreateLaporan) // hanya superadmin, programmer & level 2
	r.PUT("/laporan/:id", controllers.UpdateLaporan) // hanya superadmin, programmer & level 2
	r.DELETE("/laporan/:id", controllers.DeleteLaporan) // hanya superadmin, programmer & level 2

	r.GET("/verifikasi", controllers.GetVerifikasi) // semua users dapat mengakses
	r.POST("/verifikasi", controllers.Verifikasi) // hanya level 2
}
