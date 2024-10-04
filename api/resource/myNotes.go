package resource

import "github.com/gin-gonic/gin"

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