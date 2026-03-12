package middle_ware

import (
	"aplikasi-internal/internal/exception"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			roleInterface := c.Get("name")

			if roleInterface == nil {
				return exception.Unauthentication("role tidak ditemukan")
				// return c.JSON(http.StatusUnauthorized, "role tidak ditemukan")
			}

			roleName := roleInterface.(string)

			if roleName == "super_admin" {
				return next(c)
			}

			for _, role := range allowedRoles {
				if roleName == role {
					return next(c)
				}
			}

			return exception.AccessDenied("akses ditolak")
			// return c.JSON(http.StatusForbidden, "akses ditolak")
		}
	}
}