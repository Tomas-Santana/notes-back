package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRoutes(r *gin.RouterGroup, db *mongo.Database) {
	// create routes for all models
	CreateModelRoutes(Note{}, r, db)
}