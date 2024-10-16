package resource

import (
	"github.com/gin-gonic/gin"
)

func(rg *ResourceGroup) DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")

	if categoryId == "" {
		c.JSON(400, gin.H{"error": "Category ID is required"})
	return
	}

	err := rg.db.DeleteCategory(categoryId)

	if err!= nil {
	c.JSON(500, gin.H{"error": err.Error()})
	return
  }

	c.JSON(200, gin.H{
		"_id": categoryId,
	})

}