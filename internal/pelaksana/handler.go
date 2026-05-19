package pelaksana

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
			return exception.BadRequest("distribusi_id atau programmer_id tidak ditemukan")
		case "23505":
			return exception.Conflict("Kombinasi distribusi dan programmer ini sudah ada")
		}
	}
	return exception.InternalServer("Terjadi kesalahan pada server")
}

// GetPelaksana godoc
// @Summary Ambil semua Pelaksana
// @Description Mendapatkan daftar pelaksana
// @Tags Pelaksana
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]pelaksana.PelaksanaDetailResponse}
// @Failure 500 {object} helpers.APIResponse
// @Router /pelaksana [get]
func GetPelaksana(c echo.Context) error {
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	result, err := GetPelaksanaDetailServices(userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPelaksanaID godoc
// @Summary Ambil pelaksana berdasarkan ID
// @Description Mendapatkan data pelaksana berdasarkan UUID
// @Tags Pelaksana
// @Produce json
// @Param id path string true "Pelaksana ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=pelaksana.PelaksanaDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /pelaksana/{id} [get]
func GetPelaksanaID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetPelaksanaDetailServicesID(id)
	if err != nil {
		return handleDBError(err)
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
// @Success 201 {object} helpers.APIResponse{data=pelaksana.PelaksanaDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /pelaksana [post]
func CreatePelaksana(c echo.Context) error {
	var req PelaksanaRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := CreatePelaksanaServices(req.DistribusiID, req.ProgrammerID)
	if err != nil {
		return handleDBError(err)
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
// @Success 200 {object} helpers.APIResponse{data=pelaksana.PelaksanaDetailResponse}
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
		return handleDBError(err)
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
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}

// MarkAllReadPelaksana godoc
// @Summary Tandai semua pelaksana sebagai dibaca
// @Description Menandai semua penugasan milik programmer yang sedang login sebagai sudah dibaca
// @Tags Pelaksana
// @Produce json
// @Success 200 {object} helpers.APIResponse
// @Failure 401 {object} helpers.APIResponse
// @Failure 500 {object} helpers.APIResponse
// @Router /pelaksana/mark-all-read [patch]
func MarkAllReadPelaksana(c echo.Context) error {
	userIDInterface := c.Get("user_id")
	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)
	if err := MarkAllReadPelaksanaServices(userID); err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menandai semua penugasan sebagai dibaca", nil))
}
