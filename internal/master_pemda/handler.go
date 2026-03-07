package master_pemda

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetPemda godoc
// @Summary Ambil semua master pemda
// @Description Mendapatkan daftar master pemda
// @Tags Master Pemda
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]MasterPemda}
// @Failure 500 {object} helpers.APIResponse
// @Router /master-pemda [get]
func GetPemda(c echo.Context) error {
	result, err := GetMasterPemdaServices()

	if err != nil {
        return err
    }

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPemdaID godoc
// @Summary Ambil master pemda berdasarkan ID
// @Description Mendapatkan data master pemda berdasarkan UUID
// @Tags Master Pemda
// @Produce json
// @Param id path string true "Master Pemda ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]MasterPemda}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /master-pemda/{id} [get]
func GetPemdaID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	result, err := GetMasterPemdaServicesID(id)

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

// CreatePemda godoc
// @Summary Membuat master pemda baru
// @Description Menambahkan data master pemda
// @Tags Master Pemda
// @Accept json
// @Produce json
// @Param request body MasterPemdaRequest true "Data master pemda"
// @Success 201 {object} helpers.APIResponse{data=MasterPemda}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /master-pemda [post]
func CreatePemda(c echo.Context) error {
	var req MasterPemdaRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, 
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	result, err := CreateMasterPemdaServices(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, 
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdatePemda godoc
// @Summary Update master pemda
// @Description Mengupdate data master pemda
// @Tags Master Pemda
// @Accept json
// @Produce json
// @Param id path string true "ID Master Pemda"
// @Param request body MasterPemdaRequest true "Data master pemda"
// @Success 200 {object} helpers.APIResponse{data=MasterPemda}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /master-pemda/{id} [put]
func UpdatePemda(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	var req MasterPemdaRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	result, err := UpdateMasterPemdaServices(id, req)
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

// DeletePemda godoc
// @Summary Hapus master pemda
// @Description Menghapus data master pemda berdasarkan ID
// @Tags Master Pemda
// @Produce json
// @Param id path string true "ID Master Pemda"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /master-pemda/{id} [delete]
func DeletePemda(c echo.Context) error {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	err := DeleteMasterPemdaServices(id)
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