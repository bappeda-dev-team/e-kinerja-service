package roles

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateRole godoc
// @Summary Buat role baru
// @Description Menambahkan role baru
// @Tags Roles
// @Accept json
// @Produce json
// @Param request body RoleRequest true "Data role"
// @Success 201 {object} helpers.APIResponse{data=Roles}
// @Failure 400 {object} helpers.APIResponse
// @Router /roles [post]
func CreateRole(c echo.Context) error {
	var req RoleRequest
	if err := c.Bind(&req); err != nil || req.Name == "" {
		return exception.BadRequest("Field 'name' harus diisi")
	}

	result, err := CreateRoleServices(req.Name, req.Description)
	if err != nil {
		return exception.InternalServer("Terjadi kesalahan pada server")
	}

	return c.JSON(http.StatusCreated, helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdateRole godoc
// @Summary Update role
// @Description Mengupdate name dan description role berdasarkan ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID (UUID)"
// @Param request body RoleRequest true "Data role"
// @Success 200 {object} helpers.APIResponse{data=Roles}
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /roles/{id} [patch]
func UpdateRole(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req RoleRequest
	if err := c.Bind(&req); err != nil || req.Name == "" {
		return exception.BadRequest("Field 'name' harus diisi")
	}

	result, err := UpdateRoleServices(id, req.Name, req.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}
		return exception.InternalServer("Terjadi kesalahan pada server")
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengupdate data", result))
}

// DeleteRole godoc
// @Summary Hapus role
// @Description Menghapus role berdasarkan ID
// @Tags Roles
// @Produce json
// @Param id path string true "Role ID (UUID)"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /roles/{id} [delete]
func DeleteRole(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeleteRoleServices(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}
		return exception.InternalServer("Terjadi kesalahan pada server")
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}

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
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetRoleServicesID(id)

	if err != nil {
		// jika ID tidak ditemukan
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}
