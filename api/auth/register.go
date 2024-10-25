package auth

import (
	"fmt"
	"notes-back/controllers/auth"
	"notes-back/helpers"
	"notes-back/types"

	"notes-back/types/requestTypes"

	"strings"

	"github.com/gin-gonic/gin"
)

func (a *AuthRouter) Register(c *gin.Context) {
	var payload requestTypes.Register
	if err := helpers.ValidatePayload(c, a.validator, &payload); err != nil {
		return
	}

	user := types.User{
		Email:     strings.ToLower(payload.Email),
		Password:  auth.HashPassword(payload.Password),
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	err := a.db.CreateUser(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token := auth.CreateToken(&user)

	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})

}


func (a *AuthRouter) VerifyEmailAvialability(c *gin.Context) {
	var payload requestTypes.VerifyEmail
	if err := helpers.ValidatePayload(c, a.validator, &payload); err != nil {
		return
	}
	_, err := a.db.GetUserByEmail(
		strings.ToLower(payload.Email),
	)

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"available": true,
		})
		return
	}

	c.JSON(200, gin.H{
		"available": false,
	})
}