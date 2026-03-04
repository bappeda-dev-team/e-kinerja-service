package user

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Auth(c *gin.Context) {

}

func GetAllUser(c *gin.Context) {
	result, err := GetUserServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

func GetUserID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
		return
	}

	result, err := GetUserServicesID(id)

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

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
		return
	}

	user, err := RegisterUserService(req)
	if err != nil {

		if err.Error() == "role tidak ditemukan" ||
			err.Error() == "username sudah digunakan" {

			c.JSON(http.StatusBadRequest,
				helpers.ErrorResponse(400, err.Error(), nil))
			return
		}

		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Registrasi berhasil", user))
}

func Logout(c *gin.Context) {

}