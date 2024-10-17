package auth

import (
	"github.com/gin-gonic/gin"
	"notes-back/types/requestTypes"
	"notes-back/helpers"
	"notes-back/controllers/auth"
	"github.com/resend/resend-go/v2"
	"notes-back/controllers/email"
)

func (a *AuthRouter) SendResetEmail(c *gin.Context) {
	var payload requestTypes.SendResetEmail

	
	if err := helpers.ValidatePayload(c, a.validator, &payload); err != nil {
		return
	}
	user, err := a.db.GetUserByEmail(payload.Email)

	if err != nil {
		c.JSON(500, gin.H{"error": "Email no encontrado"})
		return
	}
	
	code := auth.CreateResetCode()

	a.resetCodes = append(a.resetCodes, ResetCode{
		Code:  code,
		Email: payload.Email,
	})

	params := &resend.SendEmailRequest{
        To:      []string{payload.Email},
        From:    "no-reply@notebit.cervant.chat",
        Html:    email.PasswordResetTemplate(code, user.FirstName),
        Subject: "Cambia tu contraseña de notebit",
    }

	_, err = a.emailClient.Emails.Send(params)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Email enviado",
	})

}