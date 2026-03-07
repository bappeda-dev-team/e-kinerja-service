package permintaan

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetPermintaan godoc
// @Summary Ambil semua permintaan
// @Description Mendapatkan daftar permintaan
// @Tags Permintaan
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]Permintaan}
// @Failure 500 {object} helpers.APIResponse
// @Router /permintaan [get]
func GetPermintaan(c echo.Context) error {
	result, err := GetPermintaanServices()

	if err != nil {
        return err
    }

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPermintaanId godoc
// @Summary Ambil permintaan berdasarkan ID
// @Description Mendapatkan data permintaan berdasarkan UUID
// @Tags Permintaan
// @Produce json
// @Param id path string true "Permitaan ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]Permintaan}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /permintaan/{id} [get]
func GetPermintaanId(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	result, err := GetPermintaanServicesID(id)

	if err != nil {

		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPermintaanNama godoc
// @Summary Ambil semua permintaan yang sudah tertampil nama
// @Description Mendapatkan daftar permintaan yang sudah tertampil nama
// @Tags Permintaan
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]PermintaanByNama}
// @Failure 500 {object} helpers.APIResponse
// @Router /permintaan-nama [get]
func GetPermintaanNama(c echo.Context) error {
	result, err := GetPermintaanNamaServices()

	if err != nil {
        return err
    }

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPermintaanNamaId godoc
// @Summary Ambil permintaan yang sudah tertampil nama berdasarkan ID
// @Description Mendapatkan data permintaan yang sudah tertampil nama berdasarkan UUID
// @Tags Permintaan
// @Produce json
// @Param id path string true "Permitaan ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]PermintaanByNama}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /permintaan-nama/{id} [get]
func GetPermintaanNamaId(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	result, err := GetPermintaanNamaServicesID(id)

	if err != nil {

		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// CreatePermintaan godoc
// @Summary Membuat permintaan baru
// @Description Menambahkan data permintaan
// @Tags Permintaan
// @Accept json
// @Produce json
// @Param request body PermintaanRequest true "Data permintaan"
// @Success 201 {object} helpers.APIResponse{data=Permintaan}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /permintaan [post]
func CreatePermintaan(c echo.Context) error {
	var req PermintaanRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, 
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	result, err := CreatePermintaanServices(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, 
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdatePermintaan godoc
// @Summary Update permintaan
// @Description Mengupdate data permintaan
// @Tags Permintaan
// @Accept json
// @Produce json
// @Param id path string true "ID Permintaan"
// @Param request body PermintaanRequest true "Data permintaan"
// @Success 200 {object} helpers.APIResponse{data=Permintaan}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /permintaan/{id} [put]
func UpdatePermintaan(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	var req PermintaanRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	result, err := UpdatePermintaanServices(id, req)
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

// DeletePermintaan godoc
// @Summary Hapus permintaan
// @Description Menghapus data permintaan berdasarkan ID
// @Tags Permintaan
// @Produce json
// @Param id path string true "ID Permintaan"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /permintaan/{id} [delete]
func DeletePermintaan(c echo.Context) error {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	err := DeletePermintaanServices(id)
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