package resource

import (
	"github.com/gin-gonic/gin"
	"notes-back/types/requestTypes"
	"notes-back/helpers"
)

func (rg *ResourceGroup) PutUser(c *gin.Context) {
	var req requestTypes.UpdateUser
	

	if err := helpers.ValidatePayload(c, rg.validator, &req); err != nil {
		return
	}

	userID, ok := c.Get("userID")

	if !ok {
		c.JSON(403, gin.H{"error": "user not found"})
		return
	}

	if req.ID != userID {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	err := rg.db.UpdateUser(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{
		"_id":       req.ID,
		"firstName": req.FirstName,
		"lastName":  req.LastName,
	})
}
