package penilaian

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreatePenilaian(c echo.Context) error {
	penilaiID := c.Get("user_id").(string)

	var req CreatePenilaianRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := CreatePenilaianService(req, penilaiID)
	if err != nil {
		msg := err.Error()
		if msg == "distribusi tidak ditemukan" ||
			msg == "tingkat_keberhasilan harus antara 0 dan 100" ||
			msg == "format tanggal_selesai tidak valid, gunakan ISO 8601" {
			return exception.BadRequest(msg)
		}
		return exception.InternalServer("Terjadi kesalahan pada server")
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Penilaian berhasil dibuat", result))
}

func UpdatePenilaian(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req UpdatePenilaianRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := UpdatePenilaianService(id, req)
	if err != nil {
		if err.Error() == "penilaian tidak ditemukan" {
			return exception.ResourceNotFound("Penilaian tidak ditemukan")
		}
		if err.Error() == "format tanggal_selesai tidak valid" {
			return exception.BadRequest(err.Error())
		}
		return exception.InternalServer("Terjadi kesalahan pada server")
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Penilaian berhasil diupdate", result))
}

func GetAllPenilaian(c echo.Context) error {
	result, err := GetAllPenilaianService()
	if err != nil {
		return exception.InternalServer("Terjadi kesalahan pada server")
	}
	if result == nil {
		result = []PenilaianResponse{}
	}
	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengambil data penilaian", result))
}

func GetPenilaianByDistribusi(c echo.Context) error {
	distribusiID := c.Param("distribusi_id")
	if _, err := uuid.Parse(distribusiID); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetByDistribusiIDService(distribusiID)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Penilaian tidak ditemukan")
		}
		return exception.InternalServer("Terjadi kesalahan pada server")
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengambil data penilaian", result))
}
