package laporan

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetLaporan godoc
// @Summary Ambil semua Laporan
// @Description Mendapatkan daftar laporan
// @Tags Laporan
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]Laporan}
// @Failure 500 {object} helpers.APIResponse
// @Router /laporan [get]
func GetLaporan(c *gin.Context) {
	result, err := GetLaporanServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
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
func GetLaporanID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	result, err := GetLaporanServicesID(id)

	if err != nil {

		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
			return
		}

		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
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
func CreateLaporan(c *gin.Context) {
	var req LaporanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
		return
	}

	result, err := CreateLaporanServices(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, 
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
func UpdateLaporan(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	var req LaporanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
		return
	}

	result, err := UpdateLaporanServices(id, req)
	if err != nil {

		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
			return
		}

		c.Error(err)
		return
	}

	c.JSON(http.StatusOK,
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
func DeleteLaporan(c *gin.Context) {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	err := DeleteLaporanServices(id)
	if err != nil {

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
			return
		}

		c.Error(err)
		return
	}

	c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}