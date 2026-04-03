package laporan

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
			return exception.BadRequest("permintaan_id tidak ditemukan")
		case "23505":
			return exception.Conflict("Data laporan untuk permintaan ini sudah ada")
		}
	}
	return exception.InternalServer("Terjadi kesalahan pada server")
}

// GetLaporan godoc
// @Summary Ambil semua Laporan
// @Description Mendapatkan daftar laporan
// @Tags Laporan
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]LaporanDetailResponse}
// @Failure 500 {object} helpers.APIResponse
// @Router /laporan [get]
func GetLaporan(c echo.Context) error {
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
// @Success 200 {object} helpers.APIResponse{data=LaporanDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /laporan/{id} [get]
func GetLaporanID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetLaporanServicesID(id)
	if err != nil {
		return handleDBError(err)
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
// @Success 201 {object} helpers.APIResponse{data=LaporanDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /laporan [post]
func CreateLaporan(c echo.Context) error {
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	var req LaporanRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := CreateLaporanServices(req.PermintaanID, userID, req)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}
func CreateVerif(c echo.Context) error {
	userIDInterface := c.Get("user_id")
	LaporanID := c.Param("laporan_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(LaporanID); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := CreateVerifikasiService(userID, LaporanID)
	if err != nil {
		return handleDBError(err)
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
// @Success 200 {object} helpers.APIResponse{data=LaporanDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /laporan/{id} [put]
func UpdateLaporan(c echo.Context) error {
	id := c.Param("id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req LaporanRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := UpdateLaporanServices(id, req.PermintaanID, userID, req)
	if err != nil {
		return handleDBError(err)
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

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeleteLaporanServices(id)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}



