package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte("secret-key")

type JwtCustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	RoleID   string `json:"role_id"`
	RoleName string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, username string, roleID string, roleName string) (string, error) {

	claims := JwtCustomClaims{
		UserID:   userID,
		Username: username,
		RoleID:   roleID,
		RoleName:   roleName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SECRET_KEY)
}