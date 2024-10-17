package types

import "github.com/golang-jwt/jwt/v5"


type UserClaims struct {
	UserID string `json:"userID"`
	Firstname string `json:"name"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}