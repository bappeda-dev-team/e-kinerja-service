package verifikasi

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetVerifikasi godoc
// @Summary Ambil semua Verifikasi
// @Description Mendapatkan daftar verifikasi
// @Tags Verifikasi
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]Verifikasi}
// @Failure 500 {object} helpers.APIResponse
// @Router /verifkasi [get]
func GetVerifikasi(c echo.Context) error {
	result, err := GetVerifikasiServices()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetVerifikasiID godoc
// @Summary Ambil verifkasi berdasarkan ID
// @Description Mendapatkan data verifikasi berdasarkan UUID
// @Tags Verifikasi
// @Produce json
// @Param id path string true "Verifikasi ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]Verifikasi}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /verifikasi/{id} [get]
func GetVerifikasiID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetVerifikasiServicesID(id)

	if err != nil {
		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// CreateVerifikasi godoc
// @Summary Membuat verifikasi baru
// @Description Menambahkan data verifikasi
// @Tags Verifikasi
// @Accept json
// @Produce json
// @Param request body VerifikasiRequest true "Data Verifikasi"
// @Success 201 {object} helpers.APIResponse{data=Verifikasi}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /verifikasi [post]
func CreateVerifikasi(c echo.Context) error {
	laporanID := c.Param("laporan_id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(laporanID); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "laporan_id tidak valid", nil))
	}

	var req VerifikasiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := CreateVerifikasiServices(laporanID, userID, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdateVerifikasi godoc
// @Summary Update verifikasi
// @Description Mengupdate data verifikasi
// @Tags Verifikasi
// @Accept json
// @Produce json
// @Param id path string true "ID Verifikasi"
// @Param request body VerifikasiRequest true "Data verifikasi"
// @Success 200 {object} helpers.APIResponse{data=Verifikasi}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /verifikasi/{id} [put]
func UpdateVerifikasi(c echo.Context) error {
	id := c.Param("id")
	laporanID := c.Param("laporan_id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	if _, err := uuid.Parse(laporanID); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "laporan_id tidak valid", nil))
	}

	var req VerifikasiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := UpdateVerifikasiServices(id, laporanID, userID, req)
	if err != nil {
		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengupdate data", result))
}

// DeleteVerifikasi godoc
// @Summary Hapus verifikasi
// @Description Menghapus data verifikasi berdasarkan ID
// @Tags Verifikasi
// @Produce json
// @Param id path string true "ID Verifikasi"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /verifikasi/{id} [delete]
func DeleteVerifikasi(c echo.Context) error {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeleteVerifikasiServices(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}
