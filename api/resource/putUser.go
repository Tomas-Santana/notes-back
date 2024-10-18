package resource

import (

	"notes-back/types/requestTypes"
	"github.com/gin-gonic/gin"
)

func(rg *ResourceGroup) PutUser(c *gin.Context) {
	var req requestTypes.UpdateUser
  if err := c.ShouldBindJSON(&req); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

	err := rg.db.UpdateUser(&req)
	if err!= nil {
    c.JSON(500, gin.H{"error": "Failed to update user"})
    return
  }

	c.JSON(200, gin.H{
		"user": req.ID,
	})
}