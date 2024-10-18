package resource

import (
	"github.com/gin-gonic/gin"
)

func(rg *ResourceGroup) DeleteUser(c *gin.Context) {
	userId := c.Param("id")

  if userId == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	err := rg.db.DeleteUser(userId)

	if err!= nil {
    c.JSON(500, gin.H{"error": err.Error()})
    return
  }

	c.JSON(200, gin.H{"message": "User deleted successfully"})

}