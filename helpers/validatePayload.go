package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidatePayload[T any](c *gin.Context, validator *validator.Validate, payload *T) error {
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return err
	}

	err = validator.Struct(payload)
	if err != nil {
		c.JSON(400, gin.H{"error": ParseError(err)})
		return err
	}
	return nil
}
