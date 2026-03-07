package distribusi

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetDistribusi godoc
// @Summary Ambil semua Distribusi
// @Description Mendapatkan daftar distribusi
// @Tags Distribusi
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]Distribusi}
// @Failure 500 {object} helpers.APIResponse
// @Router /distribusi [get]
func GetDistribusi(c echo.Context) error {
	result, err := GetDistribusiServices()

	if err != nil {
        return err
    }

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetDsitribusiById godoc
// @Summary Ambil distribusi berdasarkan ID
// @Description Mendapatkan data distribusi berdasarkan UUID
// @Tags Distribusi
// @Produce json
// @Param id path string true "Distribusi ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]Distribusi}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi/{id} [get]
func GetDistribusiById(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	result, err := GetDistribusiServicesID(id)

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

// GetDistribusiByNama godoc
// @Summary Ambil semua distribusi yang sudah tertampil nama
// @Description Mendapatkan daftar distribusi yang sudah tertampil nama
// @Tags Distribusi
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]DistribusiByNama}
// @Failure 500 {object} helpers.APIResponse
// @Router /distribusi-nama [get]
func GetDistribusiByNama(c echo.Context) error {
	result, err := GetDistribusiNamaServices()

	if err != nil {
        return err
    }

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetDistribusiByNamaId godoc
// @Summary Ambil distribusi yang sudah tertampil nama berdasarkan ID
// @Description Mendapatkan data distribusi yang sudah tertampil nama berdasarkan UUID
// @Tags Distribusi
// @Produce json
// @Param id path string true "Distribusi ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]DistribusiByNama}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi-nama/{id} [get]
func GetDistribusiByNamaId(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	result, err := GetDistribusiNamaServicesID(id)

	if err != nil {
		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengupdate data", result))
}

// CreateDistribusi godoc
// @Summary Membuat distribusi baru
// @Description Menambahkan data distribusi
// @Tags Distribusi
// @Accept json
// @Produce json
// @Param request body DistribusiRequest true "Data distribusi"
// @Success 201 {object} helpers.APIResponse{data=Distribusi}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi [post]
func CreateDistribusi(c echo.Context) error {
	var req DistribusiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, 
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	result, err := CreateDistribusiServices(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, 
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdateDistribusi godoc
// @Summary Update distribusi
// @Description Mengupdate data distribusi
// @Tags Distribusi
// @Accept json
// @Produce json
// @Param id path string true "ID Distribusi"
// @Param request body DistribusiRequest true "Data distribusi"
// @Success 200 {object} helpers.APIResponse{data=Distribusi}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi/{id} [put]
func UpdateDistribusi(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	var req DistribusiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	result, err := UpdateDistribusiServices(id, req)
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

// DeleteDistribusi godoc
// @Summary Hapus distribusi
// @Description Menghapus data distribusi berdasarkan ID
// @Tags Distribusi
// @Produce json
// @Param id path string true "ID Distribusi"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /distribusi/{id} [delete]
func DeleteDistribusi(c echo.Context) error {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	err := DeleteDistribusiServices(id)
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