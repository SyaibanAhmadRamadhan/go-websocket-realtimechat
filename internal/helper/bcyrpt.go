package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func Hashing(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(fromPassword), nil
}

func Compare(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
