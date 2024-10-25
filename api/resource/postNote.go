package resource

import (
	"fmt"
	"notes-back/helpers"
	"notes-back/types"
	"notes-back/types/requestTypes"
	"time"
	"github.com/gin-gonic/gin"
)

func (rg *ResourceGroup) PostNote(c *gin.Context) {
	var newNote requestTypes.CreateNote

	// get userID from content
	userID, ok := c.Get("userID")
	
	if newNote.ID != "" {
		c.JSON(400, gin.H{"error": "No se puede crear una nota con un ID. Debes usar PUT /note"})
		return
	}
	
	if !ok {
		c.JSON(403, gin.H{"error": "user not found"})
		return
	}
	
	if err := helpers.ValidatePayload(c, rg.validator, &newNote); err != nil {
		return
	}

	note := types.Note{
		Title:     newNote.Title,
		Content:   newNote.Content,
		Html:      newNote.Html,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Importance: newNote.Importance,
		Categories: newNote.Categories,
	}

	fmt.Println("note", note)

	noteId, err := rg.db.CreateNote(userID.(string), &note)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"noteId": noteId,
	})

}
