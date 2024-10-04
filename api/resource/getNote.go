package resource

import (

	"github.com/gin-gonic/gin"
)

func (rg *ResourceGroup) GetNote(c *gin.Context) {
	userID, ok := c.Get("userID")

	noteId := c.Param("id")
	if noteId == "" {
		c.JSON(400, gin.H{"error": "note id is required"})
		return
	}


	if !ok {
		c.JSON(403, gin.H{"error": "user not found"})
		return
	}

	note, err := rg.db.GetNoteById(noteId)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	convertedUserId, err := rg.db.StringToId(userID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}


	if note.UserID != convertedUserId {
		
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	c.JSON(200, gin.H{
		"note": note,
	})

	
}