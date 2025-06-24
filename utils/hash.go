package utils

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	password = strings.TrimSpace(password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	password = strings.TrimSpace(password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
