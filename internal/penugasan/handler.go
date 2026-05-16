package penugasan

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func handleDBError(err error) error {
	if err == sql.ErrNoRows {
		return exception.ResourceNotFound("Data tidak ditemukan")
	}
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "23503":
			return exception.BadRequest("distribusi_pelaksana_id tidak ditemukan")
		case "23505":
			return exception.Conflict("Data sudah ada")
		}
	}
	return exception.InternalServer("Terjadi kesalahan pada server")
}

func GetAllPenugasan(c echo.Context) error {
	pelaksanaID := c.QueryParam("pelaksana_id")
	distribusiID := c.QueryParam("distribusi_id")

	userIDInterface := c.Get("user_id")
	roleInterface := c.Get("name")

	if userIDInterface == nil || roleInterface == nil {
		return exception.Unauthentication("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)
	roleName := roleInterface.(string)

	var programmerID string
	if roleName == "programmer" && pelaksanaID == "" && distribusiID == "" {
		programmerID = userID
	}

	result, err := GetAllService(pelaksanaID, distribusiID, programmerID)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

func GetPenugasanByID(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetByIDService(id)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

func CreatePenugasan(c echo.Context) error {
	var req CreatePenugasanRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := CreateService(req)
	if err != nil {
		if err.Error() != "" && err != sql.ErrNoRows {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23503" {
					return exception.BadRequest("distribusi_pelaksana_id tidak ditemukan")
				}
			}
			if err.Error()[:6] == "format" {
				return exception.BadRequest(err.Error())
			}
		}
		return handleDBError(err)
	}

	return c.JSON(http.StatusCreated, helpers.SuccessResponse(201, "Berhasil membuat penugasan", result))
}

func UpdatePenugasan(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req UpdatePenugasanRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := UpdateService(id, req)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengupdate penugasan", result))
}

func UpdateStatusPenugasan(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req UpdateStatusRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := UpdateStatusService(id, req.Status)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengupdate status penugasan", result))
}

func ReassignPenugasan(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req ReassignRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	result, err := ReassignService(id, req.DistribusiPelaksanaID)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil memindahkan penugasan", result))
}

func DeletePenugasan(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	if err := DeleteService(id); err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil menghapus penugasan", nil))
}
