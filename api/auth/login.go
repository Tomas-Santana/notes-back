package auth

import (
	"notes-back/controllers/auth"
	"notes-back/helpers"

	"notes-back/types/requestTypes"

	"github.com/gin-gonic/gin"
)


func (a *AuthRouter) Login(c *gin.Context) {
	var payload requestTypes.Login
	if err := helpers.ValidatePayload(c, a.validator, &payload); err != nil {
		return	
	}

	user, err := a.db.GetUserByEmail(payload.Email)

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid email or password"})
		return
	}

	if !auth.CheckPassword(payload.Password, user.Password) {
		c.JSON(400, gin.H{"error": "invalid email or password"})
		return
	}



	user.Password = ""

	token := auth.CreateToken(user.ID)

	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})

	
}