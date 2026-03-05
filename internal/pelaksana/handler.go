package pelaksana

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetPelaksana godoc
// @Summary Ambil semua Pelaksana
// @Description Mendapatkan daftar pelaksana
// @Tags Pelaksana
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]Pelaksana}
// @Failure 500 {object} helpers.APIResponse
// @Router /pelaksana [get]
func GetPelaksana(c *gin.Context) {
	result, err := GetPelaksanaServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPelaksanaID godoc
// @Summary Ambil pelaksana berdasarkan ID
// @Description Mendapatkan data pelaksana berdasarkan UUID
// @Tags Pelaksana
// @Produce json
// @Param id path string true "Pelaksana ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]Pelaksana}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /pelaksana/{id} [get]
func GetPelaksanaID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	result, err := GetPelaksanaServicesID(id)

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

// GetPelaksanaByNama godoc
// @Summary Ambil semua pelaksana yang sudah tertampil nama
// @Description Mendapatkan daftar pelaksana yang sudah tertampil nama
// @Tags Pelaksana
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]PelaksanaNama}
// @Failure 500 {object} helpers.APIResponse
// @Router /pelaksana-nama [get]
func GetPelaksanaByNama(c *gin.Context) {
	result, err := GetPelaksanaNamaServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPelaksanaByNamaID godoc
// @Summary Ambil pelaksana yang sudah tertampil nama berdasarkan ID
// @Description Mendapatkan data pelaksana yang sudah tertampil nama berdasarkan UUID
// @Tags Pelaksana
// @Produce json
// @Param id path string true "Pelaksana ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]PelaksanaNama}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /pelaksana-nama/{id} [get]
func GetPelaksanaByNamaID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	result, err := GetPelaksanaNamaServicesID(id)

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

// CreatePelaksana godoc
// @Summary Membuat pelaksana baru
// @Description Menambahkan data pelaksana
// @Tags Pelaksana
// @Accept json
// @Produce json
// @Param request body PelaksanaRequest true "Data pelaksana"
// @Success 201 {object} helpers.APIResponse{data=Pelaksana}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /pelaksana [post]
func CreatePelaksana(c *gin.Context) {
	var req PelaksanaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
		return
	}

	result, err := CreatePelaksanaServices(req)
	if err != nil {
		if err.Error() == "programmer sudah ditugaskan di distribusi ini"||
			err.Error() == "programmer sudah ditugaskan di distribusi ini" {

			c.JSON(http.StatusBadRequest,
				helpers.ErrorResponse(400, err.Error(), nil))
			return
		}
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, 
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdatePelaksana godoc
// @Summary Update pelaksana
// @Description Mengupdate data pelaksana
// @Tags Pelaksana
// @Accept json
// @Produce json
// @Param id path string true "ID Pelaksana"
// @Param request body PelaksanaRequest true "Data pelaksana"
// @Success 200 {object} helpers.APIResponse{data=Pelaksana}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /pelaksana/{id} [put]
func UpdatePelaksana(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	var req PelaksanaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
		return
	}

	result, err := UpdatePelaksanaServices(id, req)
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

// DeletePelaksana godoc
// @Summary Hapus pelaksana
// @Description Menghapus data pelaksana berdasarkan ID
// @Tags Pelaksana
// @Produce json
// @Param id path string true "ID Pelaksana"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /pelaksana/{id} [delete]
func DeletePelaksana(c *gin.Context) {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	err := DeletePelaksanaServices(id)
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