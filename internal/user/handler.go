package user

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// func Auth(c echo.Context) error {

// }

func GetAllUser(c echo.Context) error {
	result, err := GetUserServices()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

func GetUserID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetUserServicesID(id)

	if err != nil {
		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

func Create(c echo.Context) error {
	var req RegisterRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	user, err := CreateUserService(req)
	if err != nil {

		if err.Error() == "role tidak ditemukan" ||
			err.Error() == "username sudah digunakan" {
			return exception.BadRequest(err.Error())
		}

		return err
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Registrasi berhasil", user))
}

// func Logout(c echo.Context) error {

// }
