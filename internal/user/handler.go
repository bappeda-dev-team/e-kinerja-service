package user

import (
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	var req LoginRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	token, err := LoginService(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized,
			helpers.ErrorResponse(401, err.Error(), nil))
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "login berhasil", map[string]string{
			"token": token,
		}))
}

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
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "UUID tidak valid", nil))
	}

	result, err := GetUserServicesID(id)

	if err != nil {

		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

func Create(c echo.Context) error {
	var req RegisterRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	user, err := CreateUserService(req)
	if err != nil {

		if err.Error() == "role tidak ditemukan" ||
			err.Error() == "username sudah digunakan" {

			return c.JSON(http.StatusBadRequest,
				helpers.ErrorResponse(400, err.Error(), nil))
		}

		return err
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Registrasi berhasil", user))
}

 func Logout(c echo.Context) error {
	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Logout berhasil", nil,))
 }