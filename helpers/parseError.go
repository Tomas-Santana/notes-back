package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ParseError(err error) []string {

	errors := []string{}
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return []string{"Invalid validation error"}
	}
	
	for _, err := range err.(validator.ValidationErrors) {
		message := fmt.Sprintf("Field %s failed on condition: %s", err.Field(), err.Tag())
		errors = append(errors, message)
	}

	return errors
}

