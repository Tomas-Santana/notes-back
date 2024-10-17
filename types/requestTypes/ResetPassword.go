package requestTypes

type ResetPassword struct {
	Code    string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}