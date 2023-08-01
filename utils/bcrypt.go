package utils

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password []byte) (string, error) {
	b, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func CompareHashAndPassword(hashedPassword, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password); err != nil {
		return false
	}
	return true
}
