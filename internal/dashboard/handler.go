package dashboard

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetDashboard godoc
// @Summary Dashboard per role
// @Description Mengambil data dashboard sesuai role yang sedang login (super_admin, admin, programmer, verifikator)
// @Tags Dashboard
// @Produce json
// @Success 200 {object} helpers.APIResponse
// @Failure 401 {object} helpers.APIResponse
// @Failure 403 {object} helpers.APIResponse
// @Failure 500 {object} helpers.APIResponse
// @Router /dashboard [get]
func GetDashboard(c echo.Context) error {
	userIDInterface := c.Get("user_id")
	roleInterface := c.Get("name")

	if userIDInterface == nil || roleInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)
	roleName := roleInterface.(string)

	switch roleName {
	case "super_admin":
		result, err := GetSuperAdminDashboardService()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data dashboard", result))

	case "admin":
		result, err := GetAdminDashboardService(userID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data dashboard", result))

	case "programmer":
		result, err := GetProgrammerDashboardService(userID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data dashboard", result))

	case "verifikator":
		result, err := GetVerifikatorDashboardService()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data dashboard", result))

	default:
		return exception.AccessDenied("akses ditolak")
	}
}

// GetActivity godoc
// @Summary Activity feed per role
// @Description Mengambil feed aktivitas sesuai role yang sedang login
// @Tags Dashboard
// @Produce json
// @Success 200 {object} helpers.APIResponse
// @Failure 401 {object} helpers.APIResponse
// @Failure 500 {object} helpers.APIResponse
// @Router /dashboard/activity [get]
func GetActivity(c echo.Context) error {
	userIDInterface := c.Get("user_id")
	roleInterface := c.Get("name")

	if userIDInterface == nil || roleInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)
	roleName := roleInterface.(string)

	switch roleName {
	case "super_admin":
		items, err := GetSuperAdminActivityService()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data aktivitas", items))

	case "admin":
		items, err := GetAdminActivityService(userID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data aktivitas", items))

	case "programmer":
		items, err := GetProgrammerActivityService(userID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data aktivitas", items))

	case "verifikator":
		items, err := GetVerifikatorActivityService(userID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data aktivitas", items))

	default:
		return exception.AccessDenied("akses ditolak")
	}
}
