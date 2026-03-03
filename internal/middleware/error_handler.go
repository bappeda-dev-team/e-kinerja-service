package middleware

import (
	"net/http"

	"aplikasi-internal/internal/helpers"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			c.JSON(http.StatusInternalServerError,
				helpers.ErrorResponse(500, "Terjadi kesalahan pada server", err.Error()))
		}
	}
}
