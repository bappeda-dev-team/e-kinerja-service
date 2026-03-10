package middle_ware

import (
	"aplikasi-internal/internal/helpers"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "token tidak ada")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			return c.JSON(http.StatusUnauthorized, "format token salah")
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(
			tokenString,
			&helpers.JwtCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return helpers.SECRET_KEY, nil
			},
		)

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "token tidak valid")
		}

		claims := token.Claims.(*helpers.JwtCustomClaims)

		// simpan ke context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role_id", claims.RoleID)
		c.Set("name", claims.RoleName)

		return next(c)
	}
}
