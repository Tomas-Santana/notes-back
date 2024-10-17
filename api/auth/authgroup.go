package auth

import (
	"notes-back/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/resend/resend-go/v2"
)

type ResetCode struct {
	Code string `json:"code" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type AuthRouter struct {
	db        database.Database
	group     *gin.RouterGroup
	validator *validator.Validate
	emailClient *resend.Client
}

func NewAuthRouter(db database.Database, group *gin.RouterGroup, validator *validator.Validate, emailClient *resend.Client) *AuthRouter {
	return &AuthRouter{
		db:        db,
		group:     group,
		validator: validator,
		emailClient: emailClient,
	}
}

func (ag *AuthRouter) RegisterRoutes() {
	ag.group.POST("/register", ag.Register)
	ag.group.POST("/login", ag.Login)
	ag.group.POST("/send-reset-email", ag.SendResetEmail)
	ag.group.POST("/verify-reset-code", ag.VerifyResetCode)
	ag.group.POST("/reset-password", ag.ResetPassword)
}
   