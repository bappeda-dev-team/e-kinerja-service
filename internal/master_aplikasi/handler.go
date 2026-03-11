package master_aplikasi

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetAplikasi godoc
// @Summary Ambil semua master aplikasi
// @Description Mendapatkan daftar master aplikasi
// @Tags Master Aplikasi
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]MasterAplikasi}
// @Failure 500 {object} helpers.APIResponse
// @Router /master-aplikasi [get]
func GetAplikasi(c echo.Context) error {
	result, err := GetMasterAplikasiServices()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
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
func GetAplikasiID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetMasterAplikasiServicesID(id)

	if err != nil {
		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
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
func CreateAplikasi(c echo.Context) error {
	var req CreateMasterAplikasiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := CreateMasterAplikasiServices(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated,
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
func UpdateAplikasi(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req CreateMasterAplikasiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := UpdateMasterAplikasiServices(id, req)
	if err != nil {
		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK,
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
func DeleteAplikasi(c echo.Context) error {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeleteMasterAplikasiServices(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}
