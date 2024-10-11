package shared

import (
	"crypto/sha256"
	"crypto/subtle"
	"golang.org/x/crypto/pbkdf2"
)

func HashPassword(password string, salt []byte) []byte {
	// Use PBKDF2 to hash the password with the salt
	hashedPassword := pbkdf2.Key([]byte(password), salt, 10000, 32, sha256.New)
	return hashedPassword
}

func VerifyPassword(password string, salt []byte, storedHash []byte) bool {
	computedHash := pbkdf2.Key([]byte(password), salt, 10000, 32, sha256.New)
	// Use constant time comparison to prevent timing attacks
	return subtle.ConstantTimeCompare(computedHash, storedHash) == 1
}
