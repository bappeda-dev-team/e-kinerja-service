package refresh_token

import (
	"aplikasi-internal/config"
	"aplikasi-internal/internal/helpers"
	"database/sql"
	"errors"
	"time"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Refresh(oldToken string) (*TokenPair, error) {
	rt, err := FindByToken(oldToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("refresh token tidak valid")
		}
		return nil, err
	}

	if time.Now().After(rt.ExpiresAt) {
		_ = DeleteByToken(oldToken)
		return nil, errors.New("refresh token sudah expired")
	}

	// ambil data user untuk generate access token baru
	var userID, username, roleID, roleName string
	err = config.DB.QueryRow(`
		SELECT u.id, u.username, u.role_id, r.name
		FROM users u
		JOIN roles r ON u.role_id = r.id
		WHERE u.id = $1 AND u.is_active = true
	`, rt.UserID).Scan(&userID, &username, &roleID, &roleName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user tidak ditemukan atau tidak aktif")
		}
		return nil, err
	}

	// rotation: hapus token lama
	if err := DeleteByToken(oldToken); err != nil {
		return nil, err
	}

	accessToken, err := helpers.GenerateToken(userID, username, roleID, roleName)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := helpers.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	if err := Save(userID, newRefreshToken, time.Now().Add(7*24*time.Hour)); err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
