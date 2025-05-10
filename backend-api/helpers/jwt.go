package helpers

import (
	"dummy-fullstack-go-react/backend-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute) // Token berlaku selama 24 jam

	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecretKey)

	return token, nil
}
