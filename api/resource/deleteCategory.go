package resource

import (
	"github.com/gin-gonic/gin"
)

func(rg *ResourceGroup) DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")

	userID, ok := c.Get("userID")

	if !ok {
		c.JSON(403, gin.H{"error": "user not found"})
		return
	}

	if categoryId == "" {
		c.JSON(400, gin.H{"error": "Category ID is required"})
	return
	}

	err := rg.db.DeleteCategory(categoryId, userID.(string))

	if err!= nil {
	c.JSON(500, gin.H{"error": err.Error()})
	return
  }

	c.JSON(200, gin.H{
		"_id": categoryId,
	})

}