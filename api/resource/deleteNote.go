package resource

import (
	"github.com/gin-gonic/gin"
)

func(rg *ResourceGroup) DeleteNoteById(c *gin.Context) {
	noteId := c.Param("id")

	if noteId == "" {
		c.JSON(400, gin.H{"error": "Note ID is required"})
    return
	}

	err := rg.db.DeleteNoteById(noteId)

	if err!= nil {
    c.JSON(500, gin.H{"error": err.Error()})
    return
  }

	c.JSON(200, gin.H{
		"_id": noteId,
	})

}