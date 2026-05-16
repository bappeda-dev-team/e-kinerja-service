package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetSecretKey() []byte {
	if key := os.Getenv("JWT_SECRET"); key != "" {
		return []byte(key)
	}
	return []byte("secret-key")
}

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
		RoleName: roleName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(GetSecretKey())
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
