package refresh_token

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type refreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func RefreshTokenHandler(c echo.Context) error {
	var req refreshRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("refresh_token wajib diisi")
	}

	pair, err := Refresh(req.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized,
			helpers.ErrorResponse(401, err.Error(), nil))
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "token berhasil diperbarui", pair))
}

func LogoutHandler(c echo.Context) error {
	userID := c.Get("user_id").(string)

	if err := DeleteByUserID(userID); err != nil {
		return exception.InternalServer("Gagal logout")
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Logout berhasil", nil))
}
