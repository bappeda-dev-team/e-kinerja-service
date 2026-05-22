package distribusi

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func handleDBError(err error) error {
	log.Printf("[handleDBError] type=%T msg=%s\n", err, err.Error())
	if err == sql.ErrNoRows {
		return exception.ResourceNotFound("Data tidak ditemukan")
	}
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "23503":
			return exception.BadRequest("permintaan_id tidak ditemukan")
		case "23505":
			return exception.Conflict("Data distribusi untuk permintaan ini sudah ada")
		}
	}
	return exception.InternalServer("Terjadi kesalahan pada server")
}

// GetDistribusi godoc
// @Summary Ambil semua Distribusi
// @Description Mendapatkan daftar distribusi
// @Tags Distribusi
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]DistribusiDetailResponse}
// @Failure 500 {object} helpers.APIResponse
// @Router /distribusi [get]
func GetDistribusi(c echo.Context) error {
	sort := c.QueryParam("sort")
	if sort == "" {
		sort = "newest"
	}
	if sort != "newest" && sort != "deadline" {
		return exception.BadRequest("sort tidak valid, gunakan: newest, deadline")
	}

	result, err := GetDistribusiDetailServices(sort)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetDistribusiById godoc
// @Summary Ambil distribusi berdasarkan ID
// @Description Mendapatkan data distribusi berdasarkan UUID
// @Tags Distribusi
// @Produce json
// @Param id path string true "Distribusi ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=DistribusiDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi/{id} [get]
func GetDistribusiById(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetDistribusiDetailServicesID(id)
	if err != nil {
		return handleDBError(err)
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
// @Success 201 {object} helpers.APIResponse{data=DistribusiDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi [post]
func CreateDistribusi(c echo.Context) error {
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	var req DistribusiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	req.NormalizePelaksana()
	if err := req.ValidatePelaksana(); err != nil {
		return exception.BadRequest(err.Error())
	}

	result, err := CreateDistribusiServices(req.PermintaanID, userID, req)
	if err != nil {
		return handleDBError(err)
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
// @Success 200 {object} helpers.APIResponse{data=DistribusiDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /distribusi/{id} [put]
func UpdateDistribusi(c echo.Context) error {
	id := c.Param("id")

	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req DistribusiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	req.NormalizePelaksana()
	if err := req.ValidatePelaksana(); err != nil {
		return exception.BadRequest(err.Error())
	}

	result, err := UpdateDistribusiServices(id, req.PermintaanID, userID, req)
	if err != nil {
		return handleDBError(err)
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
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}

func CreateKomentarDistribusi(c echo.Context) error {
	distribusiID := c.Param("distribusi_id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(distribusiID); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req KomentarDistribusiRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}


	result, err := CreateKomentarServices(distribusiID, userID, req)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23503":
				return exception.BadRequest("distribusi_id tidak ditemukan")
			}
		}

		return err
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}
