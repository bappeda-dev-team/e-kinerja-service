package pelaksana

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetPelaksana godoc
// @Summary Ambil semua Pelaksana
// @Description Mendapatkan daftar pelaksana. Gunakan ?expand=names untuk menampilkan nama lengkap.
// @Tags Pelaksana
// @Produce json
// @Param expand query string false "Gunakan 'names' untuk join nama"
// @Success 200 {object} helpers.APIResponse{data=[]PelaksanaResponse}
// @Failure 500 {object} helpers.APIResponse
// @Router /pelaksana [get]
func GetPelaksana(c echo.Context) error {
	if c.QueryParam("expand") == "names" {
		result, err := GetPelaksanaNamaServices()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
	}

	result, err := GetPelaksanaServices()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPelaksanaID godoc
// @Summary Ambil pelaksana berdasarkan ID
// @Description Mendapatkan data pelaksana berdasarkan UUID. Gunakan ?expand=names untuk menampilkan nama lengkap.
// @Tags Pelaksana
// @Produce json
// @Param id path string true "Pelaksana ID (UUID)"
// @Param expand query string false "Gunakan 'names' untuk join nama"
// @Success 200 {object} helpers.APIResponse{data=[]PelaksanaResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /pelaksana/{id} [get]
func GetPelaksanaID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	if c.QueryParam("expand") == "names" {
		result, err := GetPelaksanaNamaServicesID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound,
					helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
			}
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
	}

	result, err := GetPelaksanaServicesID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// CreatePelaksana godoc
// @Summary Membuat pelaksana baru
// @Description Menambahkan data pelaksana
// @Tags Pelaksana
// @Accept json
// @Produce json
// @Param request body PelaksanaRequest true "Data pelaksana"
// @Success 201 {object} helpers.APIResponse{data=PelaksanaResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /pelaksana [post]
func CreatePelaksana(c echo.Context) error {
	var req PelaksanaRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := CreatePelaksanaServices(req.DistribusiID, req.ProgrammerID)
	if err != nil {
		if err.Error() == "distribusi_id sudah digunakan" ||
			err.Error() == "programmer_id sudah digunakan" {

			return c.JSON(http.StatusBadRequest,
				helpers.ErrorResponse(400, err.Error(), nil))
		}

		return err
	}

	return c.JSON(http.StatusCreated,
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
// @Success 200 {object} helpers.APIResponse{data=PelaksanaResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /pelaksana/{id} [put]
func UpdatePelaksana(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req PelaksanaRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := UpdatePelaksanaServices(id, req.DistribusiID, req.ProgrammerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		if err.Error() == "distribusi_id sudah digunakan" ||
			err.Error() == "programmer_id sudah digunakan" {

			return c.JSON(http.StatusBadRequest,
				helpers.ErrorResponse(400, err.Error(), nil))
		}

		return err
	}

	return c.JSON(http.StatusOK,
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
func DeletePelaksana(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeletePelaksanaServices(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}
