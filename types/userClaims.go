package types

import "github.com/golang-jwt/jwt/v5"


type UserClaims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}