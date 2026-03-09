package laporan

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetLaporan godoc
// @Summary Ambil semua Laporan
// @Description Mendapatkan daftar laporan
// @Tags Laporan
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]Laporan}
// @Failure 500 {object} helpers.APIResponse
// @Router /laporan [get]
func GetLaporan(c echo.Context) error  {
	result, err := GetLaporanServices()

	if err != nil {
        return err
    }

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetLaporanID godoc
// @Summary Ambil laporan berdasarkan ID
// @Description Mendapatkan data laporan berdasarkan UUID
// @Tags Laporan
// @Produce json
// @Param id path string true "Laporan ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]Laporan}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /laporan/{id} [get]
func GetLaporanID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	result, err := GetLaporanServicesID(id)

	if err != nil {

		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// CreateLaporan godoc
// @Summary Membuat laporan baru
// @Description Menambahkan data laporan
// @Tags Laporan
// @Accept json
// @Produce json
// @Param request body LaporanRequest true "Data Laporan"
// @Success 201 {object} helpers.APIResponse{data=Laporan}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /laporan [post]
func CreateLaporan(c echo.Context) error {
	permintaanID := c.Param("permintaan_id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(permintaanID); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "pemda_id tidak valid", nil))
	}

	var req LaporanRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, 
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	result, err := CreateLaporanServices(permintaanID, userID,req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, 
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdateLaporan godoc
// @Summary Update laporan
// @Description Mengupdate data laporan
// @Tags Laporan
// @Accept json
// @Produce json
// @Param id path string true "ID Laporan"
// @Param request body LaporanRequest true "Data laporan"
// @Success 200 {object} helpers.APIResponse{data=Laporan}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /laporan/{id} [put]
func UpdateLaporan(c echo.Context) error {
	id := c.Param("id")
	permintaanID := c.Param("permintaan_id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	if _, err := uuid.Parse(permintaanID); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "permintaan_id tidak valid", nil))
	}

	var req LaporanRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	result, err := UpdateLaporanServices(id, permintaanID, userID, req)
	if err != nil {

		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengupdate data", result))
}

// DeleteLaporan godoc
// @Summary Hapus laporan
// @Description Menghapus data laporan berdasarkan ID
// @Tags Laporan
// @Produce json
// @Param id path string true "ID Laporan"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /laporan/{id} [delete]
func DeleteLaporan(c echo.Context) error {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	err := DeleteLaporanServices(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}