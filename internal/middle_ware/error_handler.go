package middle_ware

import (
	"net/http"

	"aplikasi-internal/internal/helpers"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {

	if c.Response().Committed {
		return
	}

	c.JSON(http.StatusInternalServerError,
		helpers.ErrorResponse(
			500,
			"Terjadi kesalahan pada server",
			err.Error(),
		))
}
