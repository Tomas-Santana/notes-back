package models

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"

	"notes-back/database"

	"strings"
)



func CreateModelRoutes(modelInstance interface{}, router *gin.RouterGroup) {

	modelName := reflect.TypeOf(modelInstance).Name()

	if modelName == "" {
		panic("could not get model name")
	}

	modelName = strings.ToLower(modelName)



	// modelSchema := getModelSchema(model)
	data := database.LoadDatabase()

	fmt.Println("Creating routes for " + modelName)
	router.GET("/" + modelName + "/:id", func(c *gin.Context) {
		
		id := c.Param("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}
		item, err := data.GetItem(modelName, intID)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, item)
	})

}

func getModelSchema[Model any](model Model) map[string]reflect.Kind {
	modelType := reflect.TypeOf(model)
	modelSchema := make(map[string]reflect.Kind)

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		fieldName := field.Name
		fieldType := field.Type.Kind()
		modelSchema[fieldName] = fieldType
	}

	return modelSchema
}
