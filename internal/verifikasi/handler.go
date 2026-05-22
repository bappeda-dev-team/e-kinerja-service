package verifikasi

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func handleDBError(err error) error {
	log.Printf("[handleDBError] type=%T msg=%s\n", err, err.Error())
	if err == sql.ErrNoRows {
		return exception.ResourceNotFound("Data tidak ditemukan")
	}
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "23503":
			return exception.BadRequest("laporan_id tidak ditemukan")
		case "23505":
			return exception.Conflict("Data verifikasi untuk laporan ini sudah ada")
		case "23514":
			return exception.BadRequest("status_verified tidak valid, gunakan: pending, approved, revision")
		}
	}
	return exception.InternalServer("Terjadi kesalahan pada server")
}

// GetVerifikasi godoc
// @Summary Ambil semua Verifikasi
// @Description Mendapatkan daftar verifikasi
// @Tags Verifikasi
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]VerifikasiDetailResponse}
// @Failure 500 {object} helpers.APIResponse
// @Router /verifikasi [get]
func GetVerifikasi(c echo.Context) error {
	result, err := GetVerifikasiDetailServices()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetVerifikasiID godoc
// @Summary Ambil verifikasi berdasarkan ID
// @Description Mendapatkan data verifikasi berdasarkan UUID
// @Tags Verifikasi
// @Produce json
// @Param id path string true "Verifikasi ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=VerifikasiDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /verifikasi/{id} [get]
func GetVerifikasiID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetVerifikasiDetailServicesID(id)
	if err != nil {
		return handleDBError(err)
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
// @Success 201 {object} helpers.APIResponse{data=VerifikasiDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /verifikasi [post]
func CreateVerifikasi(c echo.Context) error {
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)
	if userID == "" {
		return exception.Unauthorized("user_id kosong di token")
	}

	var req VerifikasiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := CreateVerifikasiServices(req.LaporanID, userID, req)
	if err != nil {
		return handleDBError(err)
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
// @Success 200 {object} helpers.APIResponse{data=VerifikasiDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /verifikasi/{id} [put]
func UpdateVerifikasi(c echo.Context) error {
	id := c.Param("id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)
	if userID == "" {
		return exception.Unauthorized("user_id kosong di token")
	}

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req VerifikasiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := UpdateVerifikasiServices(id, req.LaporanID, userID, req)
	if err != nil {
		return handleDBError(err)
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

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeleteVerifikasiServices(id)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}
