package roles

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetRoles godoc
// @Summary Ambil semua role
// @Description Mendapatkan daftar role
// @Tags Roles
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]Roles}
// @Failure 500 {object} helpers.APIResponse
// @Router /roles [get]
func GetRoles(c *gin.Context) {
	result, err := GetRolesServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))

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
func GetRoleID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	result, err := GetRoleServicesID(id)

	if err != nil {

		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
			return
		}

		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}
