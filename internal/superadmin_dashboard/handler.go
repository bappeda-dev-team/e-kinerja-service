package superadmin_dashboard

import (
	"aplikasi-internal/internal/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetSuperadminDashboard godoc
// @Summary Dashboard superadmin
// @Description Mengambil ringkasan permintaan beserta distribusi dan laporan dalam satu endpoint
// @Tags Dashboard
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=DashboardResponse}
// @Failure 500 {object} helpers.APIResponse
// @Router /superadmin-dashboard [get]
func GetSuperadminDashboard(c echo.Context) error {
	result, err := GetDashboard()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data dashboard", result))
}
