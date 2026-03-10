package middle_ware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			roleInterface := c.Get("name")

			if roleInterface == nil {
				return c.JSON(http.StatusUnauthorized, "role tidak ditemukan")
			}

			roleName := roleInterface.(string)

			for _, role := range allowedRoles {
				if roleName == role {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, "akses ditolak")
		}
	}
}