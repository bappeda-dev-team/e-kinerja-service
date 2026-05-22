package penilaian

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
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
		return err
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
		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Penilaian berhasil diupdate", result))
}

func GetAllPenilaian(c echo.Context) error {
	result, err := GetAllPenilaianService()
	if err != nil {
		return err
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
		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengambil data penilaian", result))
}
