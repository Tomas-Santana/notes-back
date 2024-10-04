package auth

import (
	"notes-back/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthRouter struct {
	db        database.Database
	group     *gin.RouterGroup
	validator *validator.Validate
}

func NewAuthRouter(db database.Database, group *gin.RouterGroup, validator *validator.Validate) *AuthRouter {
	return &AuthRouter{
		db:        db,
		group:     group,
		validator: validator,
	}
}

func (ag *AuthRouter) RegisterRoutes() {
	ag.group.POST("/register", ag.Register)
	ag.group.POST("/login", ag.Login)
}
   