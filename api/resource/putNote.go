package resource

import (
	"fmt"
	"notes-back/helpers"
	"notes-back/types"
	"notes-back/types/requestTypes"
	"time"

	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func (rg *ResourceGroup) PutNote(c *gin.Context) {
	var newNote requestTypes.NewNote

	// get userID from content
	userID, ok := c.Get("userID")

	if !ok {
		c.JSON(403, gin.H{"error": "user not found"})
		return
	}

	
	if err := helpers.ValidatePayload(c, rg.validator, &newNote); err != nil {
		return
	}

	if newNote.ID == "" {
		c.JSON(400, gin.H{"error": "Se requiere un ID"})
		return
	}

	var preview string
	if utf8.RuneCountInString(newNote.Content) < 100 {
		preview = newNote.Content
	} else {
		preview = newNote.Content[:100]
	}

	fmt.Println("newNote.ID", newNote.ID)

	note := types.Note{
		ID: newNote.ID,
		Title:   newNote.Title,
		Content: newNote.Content,
		Html: newNote.Html,
		Preview: preview,
		UpdatedAt: time.Now(),
	}



	
	err := rg.db.UpdateNote(userID.(string), &note)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"noteId": newNote.ID,
	})

		


}