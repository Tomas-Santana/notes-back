package requestTypes

type VerifyEmail struct {
	Email string `json:"email" validate:"required,email"`
}