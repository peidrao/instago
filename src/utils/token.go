package utils

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(username string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractBearerToken(headerValue string) string {
	parts := strings.Split(headerValue, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}
