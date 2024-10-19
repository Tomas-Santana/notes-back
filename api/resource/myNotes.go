package resource

import (
	"fmt"
	"notes-back/helpers"
	"notes-back/types/requestTypes"

	"github.com/gin-gonic/gin"
)

func (rg *ResourceGroup) MyNotes(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, gin.H{"error": "user not found"})
		return
	}

	notes, err := rg.db.GetUserNotes(userID.(string))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"notes": notes,
	})
}

// search notes

func (rg *ResourceGroup) SearchNotes(c *gin.Context) {
	var req requestTypes.SearchNotes

	if err := helpers.ValidatePayload(c, rg.validator, &req); err != nil {
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, gin.H{"error": "user not found"})
		return
	}

	

	notes, err := rg.db.SearchUserNotes(req.Query, userID.(string))

	fmt.Println(notes)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"notes": notes,
	})
}
