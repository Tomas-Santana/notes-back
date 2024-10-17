package auth

import (
	"notes-back/controllers/auth"
	"notes-back/controllers/email"
	"notes-back/helpers"
	"notes-back/types/requestTypes"

	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v2"
	
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

	a.db.AddResetCode(payload.Email, code)

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

func (a *AuthRouter) VerifyResetCode(c *gin.Context) {
	var payload requestTypes.VerifyResetCode

	if err := helpers.ValidatePayload(c, a.validator, &payload); err != nil {
		return
	}

	code := payload.Code

	_, err := a.db.GetResetCode(code)

	if err != nil {
		c.JSON(500, gin.H{"error": "No se encontró el código"})
	}

	c.JSON(200, gin.H{"message": "Código válido"})
}

func (a *AuthRouter) ResetPassword(c *gin.Context) {
	var payload requestTypes.ResetPassword

	if err := helpers.ValidatePayload(c, a.validator, &payload); err != nil {
		return
	}

	code := payload.Code

	email, err := a.db.GetResetCode(code)

	if err != nil {
		c.JSON(500, gin.H{"error": "No se encontró el código"})
	}

	user, err := a.db.GetUserByEmail(email)

	if err != nil {
		c.JSON(500, gin.H{"error": "No se encontró el usuario"})
	}

	user.Password = auth.HashPassword(payload.Password)

	err = a.db.UpdateUserPassword(email, user.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar la contraseña"})
	}

	c.JSON(200, gin.H{"message": "Contraseña actualizada"})

	a.db.DeleteResetCode(code)
}