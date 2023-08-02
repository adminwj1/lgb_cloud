package utils

import (
	"clouds.lgb24kcs.cn/global"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password []byte) (string, error) {
	b, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		global.APP.Log.Error(err.Error())

		return "", err
	}
	return string(b), nil
}

func CompareHashAndPassword(hashedPassword, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password); err != nil {
		global.APP.Log.Error(err.Error())

		return false
	}
	return true
}
