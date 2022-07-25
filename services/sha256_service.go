package services

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Encoder(password string) string {
	passwordHash := sha256.Sum256([]byte(password))

	// return string(passwordHash[:])
	return fmt.Sprintf("%x", passwordHash)
}
