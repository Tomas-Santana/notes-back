package resource

import (
	"notes-back/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	middleware "notes-back/middleware"
)

type ResourceGroup struct {
	db        database.Database
	group     *gin.RouterGroup
	validator *validator.Validate
}

func NewRouter(db database.Database, group *gin.RouterGroup, validator *validator.Validate) *ResourceGroup {
	return &ResourceGroup{
		db:        db,
		group:     group,
		validator: validator,
	}
}

func (rg *ResourceGroup) RegisterRoutes() {
	rg.group.Use(middleware.Authorize())
	rg.group.POST("/note", rg.PostNote)
	rg.group.GET("/notes", rg.MyNotes)
	rg.group.GET("/note/:id", rg.GetNote)
	rg.group.PUT("/note", rg.PutNote)
}
