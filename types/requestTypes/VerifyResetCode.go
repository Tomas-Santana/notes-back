package requestTypes

type VerifyResetCode struct {
	Code string `json:"code" binding:"required"`
}