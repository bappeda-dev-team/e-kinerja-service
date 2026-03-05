package master_aplikasi

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAplikasi godoc
// @Summary Ambil semua master aplikasi
// @Description Mendapatkan daftar master aplikasi
// @Tags Master Aplikasi
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]MasterAplikasi}
// @Failure 500 {object} helpers.APIResponse
// @Router /master-aplikasi [get]
func GetAplikasi(c *gin.Context) {
	result, err := GetMasterAplikasiServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetAplikasiID godoc
// @Summary Ambil master aplikasi berdasarkan ID
// @Description Mendapatkan data master aplikasi berdasarkan UUID
// @Tags Master Aplikasi
// @Produce json
// @Param id path string true "Master Aplikasi ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]MasterAplikasi}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /master-aplikasi/{id} [get]
func GetAplikasiID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	result, err := GetMasterAplikasiServicesID(id)

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

// CreateAplikasi godoc
// @Summary Membuat master aplikasi baru
// @Description Menambahkan data master aplikasi
// @Tags Master Aplikasi
// @Accept json
// @Produce json
// @Param request body CreateMasterAplikasiRequest true "Data master aplikasi"
// @Success 201 {object} helpers.APIResponse{data=MasterAplikasi}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /master-aplikasi [post]
func CreateAplikasi(c *gin.Context) {
	var req CreateMasterAplikasiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
		return
	}

	result, err := CreateMasterAplikasiServices(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, 
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdateAplikasi godoc
// @Summary Update master aplikasi
// @Description Mengupdate data master aplikasi
// @Tags Master Aplikasi
// @Accept json
// @Produce json
// @Param id path string true "ID Master Aplikasi"
// @Param request body CreateMasterAplikasiRequest true "Data master aplikasi"
// @Success 200 {object} helpers.APIResponse{data=MasterAplikasi}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /master-aplikasi/{id} [put]
func UpdateAplikasi(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	var req CreateMasterAplikasiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
		return
	}

	result, err := UpdateMasterAplikasiServices(id, req)
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

// DeleteAplikasi godoc
// @Summary Hapus master aplikasi
// @Description Menghapus data master aplikasi berdasarkan ID
// @Tags Master Aplikasi
// @Produce json
// @Param id path string true "ID Master Aplikasi"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /master-aplikasi/{id} [delete]
func DeleteAplikasi(c *gin.Context) {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	err := DeleteMasterAplikasiServices(id)
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