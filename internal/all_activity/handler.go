package all_activity

import (
	"aplikasi-internal/internal/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllActivity godoc
// @Summary Ambil semua aktivitas
// @Description Mendapatkan timeline semua aktivitas: permintaan, distribusi, pelaksana, laporan, dan verifikasi — diurutkan terbaru
// @Tags Activity
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]ActivityItem}
// @Failure 500 {object} helpers.APIResponse
// @Router /all-activity [get]
func GetAllActivity(c echo.Context) error {
	result, err := GetAllActivitiesService()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data aktivitas", result))
}
