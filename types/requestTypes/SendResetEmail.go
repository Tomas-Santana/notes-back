package requestTypes

type SendResetEmail struct {
	Email string `json:"email" binding:"required,email"`
}