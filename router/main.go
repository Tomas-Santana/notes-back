package router

import (
	"net/http"

	"github.com/gin-gonic/gin"


	"notes-back/models"
)

func Start() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	
	models.CreateModelRoutes(models.Album{}, r.Group("/api"))

	r.Run("localhost:8080")
}
