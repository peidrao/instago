package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("failed to generate random string")
	}

	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	return randomString[:length]
}
