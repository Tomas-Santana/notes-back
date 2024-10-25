package resource

import (
	"notes-back/helpers"
	"notes-back/types/requestTypes"
	"github.com/gin-gonic/gin"
)

func (rg *ResourceGroup) PutNote(c *gin.Context) {
	var noteUpdate requestTypes.UpdateNote

	userID, ok := c.Get("userID")

	if !ok {
		c.JSON(403, gin.H{"error": "user not found"})
		return
	}

	if err := helpers.ValidatePayload(c, rg.validator, &noteUpdate); err != nil {
		return
	}

	if noteUpdate.ID == "" {
		c.JSON(400, gin.H{"error": "Se requiere un ID"})
		return
	}

	err := rg.db.UpdateNote(userID.(string), &noteUpdate)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"noteId": noteUpdate.ID,
	})

}
