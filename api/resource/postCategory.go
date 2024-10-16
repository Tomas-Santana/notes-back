package resource

import (
	"notes-back/helpers"
	"notes-back/types/requestTypes"
	"notes-back/types"
	
	"github.com/gin-gonic/gin"
)

func (rg *ResourceGroup) PostCategory(c *gin.Context) {
	var newCategory requestTypes.NewCategory

	if err := helpers.ValidatePayload(c, rg.validator, &newCategory); err != nil {
		return
	}

	userID, ok := c.Get("userID")

	if !ok {
		c.JSON(403, gin.H{"error": "user not found"})
		return
	}

	category := types.Category{
		Name: newCategory.Name,
		UserID: userID.(string),
	}

	err := rg.db.CreateCategory(&category)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"category": category})
}