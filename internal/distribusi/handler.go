package distribusi

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetDistribusi godoc
// @Summary Ambil semua Distribusi
// @Description Mendapatkan daftar distribusi. Gunakan ?expand=names untuk menampilkan nama lengkap.
// @Tags Distribusi
// @Produce json
// @Param expand query string false "Gunakan 'names' untuk join nama"
// @Success 200 {object} helpers.APIResponse{data=[]DistribusiResponse}
// @Failure 500 {object} helpers.APIResponse
// @Router /distribusi [get]
func GetDistribusi(c echo.Context) error {
	if c.QueryParam("expand") == "names" {
		result, err := GetDistribusiNamaServices()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
	}

	result, err := GetDistribusiServices()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetDistribusiById godoc
// @Summary Ambil distribusi berdasarkan ID
// @Description Mendapatkan data distribusi berdasarkan UUID. Gunakan ?expand=names untuk menampilkan nama lengkap.
// @Tags Distribusi
// @Produce json
// @Param id path string true "Distribusi ID (UUID)"
// @Param expand query string false "Gunakan 'names' untuk join nama"
// @Success 200 {object} helpers.APIResponse{data=[]DistribusiResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi/{id} [get]
func GetDistribusiById(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	if c.QueryParam("expand") == "names" {
		result, err := GetDistribusiNamaServicesID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return exception.ResourceNotFound("Data tidak ditemukan")
			}
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
	}

	result, err := GetDistribusiServicesID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// CreateDistribusi godoc
// @Summary Membuat distribusi baru
// @Description Menambahkan data distribusi
// @Tags Distribusi
// @Accept json
// @Produce json
// @Param request body DistribusiRequest true "Data distribusi"
// @Success 201 {object} helpers.APIResponse{data=DistribusiResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi [post]
func CreateDistribusi(c echo.Context) error {
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	var req DistribusiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := CreateDistribusiServices(req.PermintaanID, userID, req)
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
// @Success 200 {object} helpers.APIResponse{data=DistribusiResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi/{id} [put]
func UpdateDistribusi(c echo.Context) error {
	id := c.Param("id")

	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req DistribusiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := UpdateDistribusiServices(id, req.PermintaanID, userID, req)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
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

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeleteDistribusiServices(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}
