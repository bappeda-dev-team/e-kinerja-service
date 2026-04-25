package routes

import (
	"aplikasi-internal/internal/all_activity"
	"aplikasi-internal/internal/distribusi"
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/laporan"
	"aplikasi-internal/internal/master_aplikasi"
	"aplikasi-internal/internal/master_pemda"
	"aplikasi-internal/internal/middle_ware"
	"aplikasi-internal/internal/pelaksana"
	"aplikasi-internal/internal/permintaan"
	"aplikasi-internal/internal/roles"
	"aplikasi-internal/internal/superadmin_dashboard"
	"aplikasi-internal/internal/user"
	"aplikasi-internal/internal/verifikasi"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "aplikasi-internal/docs"
)

func SetupRoutes(r *echo.Echo) {
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	// global error handler
	r.HTTPErrorHandler = exception.GlobalExceptionHandler()

	r.GET("/roles", roles.GetRoles)
	r.GET("/roles/:id", roles.GetRoleID)
	r.POST("/roles", roles.CreateRole, middle_ware.JWTMiddleware, middle_ware.RoleMiddleware("super_admin"))
	r.PATCH("/roles/:id", roles.UpdateRole, middle_ware.JWTMiddleware, middle_ware.RoleMiddleware("super_admin"))
	r.DELETE("/roles/:id", roles.DeleteRole, middle_ware.JWTMiddleware, middle_ware.RoleMiddleware("super_admin"))

	auth := r.Group("/auth")
	auth.POST("/login", user.Login)
	auth.POST("/logout", user.Logout, middle_ware.JWTMiddleware)

	r.GET("/users", user.GetAllUser)
	r.GET("/users/:id", user.GetUserID)
	r.POST("/users", user.Create, middle_ware.JWTMiddleware, middle_ware.RoleMiddleware("super_admin"))
	r.DELETE("/users/:id", user.DeleteUser, middle_ware.JWTMiddleware, middle_ware.RoleMiddleware("super_admin"))
	r.PATCH("/users/:id/profile-picture", user.UploadProfilePic, middle_ware.JWTMiddleware)
	r.PATCH("/users/:id", user.PatchUser, middle_ware.JWTMiddleware, middle_ware.RoleMiddleware("super_admin"))

	ma := r.Group("/master-aplikasi")
	ma.Use(middle_ware.JWTMiddleware)
	ma.Use(middle_ware.RoleMiddleware("super_admin"))

	ma.GET("", master_aplikasi.GetAplikasi)
	ma.GET("/:id", master_aplikasi.GetAplikasiID)
	ma.POST("", master_aplikasi.CreateAplikasi)
	ma.PUT("/:id", master_aplikasi.UpdateAplikasi)
	ma.DELETE("/:id", master_aplikasi.DeleteAplikasi)
	ma.PATCH("/:id/logo", master_aplikasi.UploadLogo)

	mp := r.Group("/master-pemda")
	mp.Use(middle_ware.JWTMiddleware)
	mp.Use(middle_ware.RoleMiddleware("super_admin"))

	mp.GET("", master_pemda.GetPemda)
	mp.GET("/:id", master_pemda.GetPemdaID)
	mp.POST("", master_pemda.CreatePemda)
	mp.PUT("/:id", master_pemda.UpdatePemda)
	mp.DELETE("/:id", master_pemda.DeletePemda)
	mp.PATCH("/:id/logo", master_pemda.UploadLogo)

	p := r.Group("/permintaan")
	p.Use(middle_ware.JWTMiddleware)
	p.Use(middle_ware.RoleMiddleware("super_admin", "admin", "programmer"))

	p.GET("", permintaan.GetPermintaan)
	p.GET("/archived", permintaan.GetArchivedPermintaan)
	p.GET("/:id", permintaan.GetPermintaanId)

	pAdmin := r.Group("/permintaan")
	pAdmin.Use(middle_ware.JWTMiddleware)
	pAdmin.Use(middle_ware.RoleMiddleware("super_admin", "admin"))

	pAdmin.POST("", permintaan.CreatePermintaan)
	pAdmin.PUT("/:id", permintaan.UpdatePermintaan)
	pAdmin.DELETE("/:id", permintaan.DeletePermintaan)
	pAdmin.PATCH("/:id/lampiran", permintaan.UploadLampiran)
	pAdmin.PATCH("/:id/status", permintaan.UpdateStatusPermintaan)

	d := r.Group("/distribusi")
	d.Use(middle_ware.JWTMiddleware)
	d.Use(middle_ware.RoleMiddleware("super_admin", "admin", "programmer"))

	d.GET("", distribusi.GetDistribusi)
	d.GET("/:id", distribusi.GetDistribusiById)
	d.POST("", distribusi.CreateDistribusi)
	d.PUT("/:id", distribusi.UpdateDistribusi)
	d.DELETE("/:id", distribusi.DeleteDistribusi)
	d.POST("/komentar/:distribusi_id", distribusi.CreateKomentarDistribusi)

	dp := r.Group("/pelaksana")
	dp.Use(middle_ware.JWTMiddleware)
	dp.Use(middle_ware.RoleMiddleware("admin", "programmer"))

	dp.GET("", pelaksana.GetPelaksana)
	dp.PATCH("/mark-all-read", pelaksana.MarkAllReadPelaksana)
	dp.GET("/:id", pelaksana.GetPelaksanaID)
	dp.POST("", pelaksana.CreatePelaksana)
	dp.PUT("/:id", pelaksana.UpdatePelaksana)
	dp.DELETE("/:id", pelaksana.DeletePelaksana)

	l := r.Group("/laporan")
	l.Use(middle_ware.JWTMiddleware)
	l.Use(middle_ware.RoleMiddleware("programmer", "verifikator"))

	l.GET("", laporan.GetLaporan)
	l.GET("/:id", laporan.GetLaporanID)
	l.GET("/history", laporan.GetHistory)
	l.POST("", laporan.CreateLaporan)
	l.POST("/verif/:laporan_id", laporan.CreateVerif)
	l.PUT("/:id", laporan.UpdateLaporan)
	l.DELETE("/:id", laporan.DeleteLaporan)
	l.PATCH("/:id/lampiran", laporan.UploadLampiran)
	l.POST("/komentar/:laporan_id", laporan.CreateKomentarLaporan)

	sd := r.Group("/superadmin-dashboard")
	sd.Use(middle_ware.JWTMiddleware)
	sd.Use(middle_ware.RoleMiddleware("super_admin"))
	sd.GET("", superadmin_dashboard.GetSuperadminDashboard)

	v := r.Group("/verifikasi")
	v.Use(middle_ware.JWTMiddleware)
	v.Use(middle_ware.RoleMiddleware("verifikator"))

	v.GET("", verifikasi.GetVerifikasi)
	v.GET("/:id", verifikasi.GetVerifikasiID)
	v.POST("", verifikasi.CreateVerifikasi)
	v.PUT("/:id", verifikasi.UpdateVerifikasi)
	v.DELETE("/:id", verifikasi.DeleteVerifikasi)

	aa := r.Group("/all-activity")
	aa.Use(middle_ware.JWTMiddleware)
	aa.Use(middle_ware.RoleMiddleware("super_admin", "admin"))
	aa.GET("", all_activity.GetAllActivity)
}
