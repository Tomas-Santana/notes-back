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
	fmt.Println("noteimportance", newNote)

	var preview string
	if utf8.RuneCountInString(newNote.Content) < 100 {
		preview = newNote.Content
	} else {
		preview = newNote.Content[:100]
	}

	noteCategories := convertToNoteCategories(newNote.Categories)

	note := types.Note{
		Title:     newNote.Title,
		Content:   newNote.Content,
		Html:      newNote.Html,
		CreatedAt: time.Now(),
		Preview:   preview,
		UpdatedAt: time.Now(),
		Importance: newNote.Importance,
		Categories: noteCategories,
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


func convertToNoteCategories(categories []string) []types.NoteCategory {
	var noteCategories []types.NoteCategory
	for _, category := range categories {
			noteCategory := types.NoteCategory{
					Name: category,
			}
			noteCategories = append(noteCategories, noteCategory)
	}
	return noteCategories
}