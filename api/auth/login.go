package auth

import (
	"fmt"
	"notes-back/controllers/auth"
	"notes-back/helpers"

	"notes-back/types/requestTypes"

	"strings"

	"github.com/gin-gonic/gin"
)

func (a *AuthRouter) Login(c *gin.Context) {
	var payload requestTypes.LoginRequest
	if err := helpers.ValidatePayload(c, a.validator, &payload); err != nil {
		return
	}

	user, err := a.db.GetUserByEmail(
		strings.ToLower(payload.Email),
	)

	if err != nil {
		fmt.Println("login error:", err)
		c.JSON(401, gin.H{"error": "Email o contraseña incorrectos"})
		return
	}

	if !auth.CheckPassword(payload.Password, user.Password) {
		c.JSON(401, gin.H{"error": "Email o contraseña incorrectos"})
		return
	}

	user.Password = ""

	token := auth.CreateToken(&user)

	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})

}


