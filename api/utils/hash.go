package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(inputPassword string) (string, error) {

	byte, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 14)

	return string(byte), err
}

func CheckPassword(inputPassword, inputHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(inputHash), []byte(inputPassword))
	return err == nil
}
