package auth

import (
	"github.com/google/uuid"
)

func CreateResetCode() string {
	return uuid.New().String()
}