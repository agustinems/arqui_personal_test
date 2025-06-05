package repositories

import (
	"time"

	"proyecto2025/backend/models"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

func SetJWTSecret(secret []byte) {
	jwtSecret = secret
}

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.Id,
		"isAdmin": user.IsAdmin,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
