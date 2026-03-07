package roles

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetRoles godoc
// @Summary Ambil semua role
// @Description Mendapatkan daftar role
// @Tags Roles
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]Roles}
// @Failure 500 {object} helpers.APIResponse
// @Router /roles [get]
func GetRoles(c echo.Context) error {

	result, err := GetRolesServices()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetRoleID godoc
// @Summary Ambil role berdasarkan ID
// @Description Mendapatkan data role berdasarkan UUID
// @Tags Roles
// @Produce json
// @Param id path string true "Role ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]Roles}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /roles/{id} [get]
func GetRoleID(c echo.Context) error {

	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	result, err := GetRoleServicesID(id)

	if err != nil {

		// jika ID tidak ditemukan
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

