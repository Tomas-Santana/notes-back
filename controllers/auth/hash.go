package auth

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	salt := os.Getenv("SALT")
	if salt == "" {
		panic("SALT must be set in .env file")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func CheckPassword(password string, hash string) bool {
	salt := os.Getenv("SALT")
	if salt == "" {
		panic("SALT must be set in .env file")
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}