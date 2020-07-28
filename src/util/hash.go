package util

import "golang.org/x/crypto/bcrypt"

func CalcHash(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
