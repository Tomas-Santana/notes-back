package auth

import (
	"notes-back/types"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"userID"`
}

func CreateToken(userId string) string {

	claims := types.UserClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
		},
	}
	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		panic(err)
	}

	return ss

}