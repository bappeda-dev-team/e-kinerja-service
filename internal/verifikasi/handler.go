package verifikasi

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetVerifikasi godoc
// @Summary Ambil semua Verifikasi
// @Description Mendapatkan daftar verifikasi
// @Tags Verifikasi
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]Verifikasi}
// @Failure 500 {object} helpers.APIResponse
// @Router /verifkasi [get]
func GetVerifikasi(c *gin.Context) {
	result, err := GetVerifikasiServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
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
func GetVerifikasiID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	result, err := GetVerifikasiServicesID(id)

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
func CreateVerifikasi(c *gin.Context) {
	var req VerifikasiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
		return
	}

	result, err := CreateVerifikasiServices(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, 
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
func UpdateVerifikasi(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	var req VerifikasiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
		return
	}

	result, err := UpdateVerifikasiServices(id, req)
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
func DeleteVerifikasi(c *gin.Context) {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	err := DeleteVerifikasiServices(id)
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