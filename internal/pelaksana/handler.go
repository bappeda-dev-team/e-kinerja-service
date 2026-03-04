package pelaksana

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, 
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

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