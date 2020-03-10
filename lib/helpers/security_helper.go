package helpers

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const hashCost = 16

// HashPassword hash and salt password with bcrypt
func HashPassword(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

// ComparePassword compare hashed password and plain password
func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	defer TimeTrack(time.Now(), "ComparePassword")

	byteHashDB := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHashDB, plainPwd)
	if err != nil {
		return false
	}

	return true
}
