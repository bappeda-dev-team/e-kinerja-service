package refresh_token

import (
	"aplikasi-internal/config"
	"time"
)

func Save(userID, token string, expiresAt time.Time) error {
	_, err := config.DB.Exec(`
		INSERT INTO refresh_tokens (user_id, token, expires_at)
		VALUES ($1, $2, $3)
	`, userID, token, expiresAt)
	return err
}

func FindByToken(token string) (*RefreshToken, error) {
	var rt RefreshToken
	err := config.DB.QueryRow(`
		SELECT id, user_id, token, expires_at, created_at
		FROM refresh_tokens
		WHERE token = $1
	`, token).Scan(&rt.ID, &rt.UserID, &rt.Token, &rt.ExpiresAt, &rt.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func DeleteByToken(token string) error {
	_, err := config.DB.Exec(`DELETE FROM refresh_tokens WHERE token = $1`, token)
	return err
}

func DeleteByUserID(userID string) error {
	_, err := config.DB.Exec(`DELETE FROM refresh_tokens WHERE user_id = $1`, userID)
	return err
}
