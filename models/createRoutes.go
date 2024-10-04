package models

import (
	"context"
	"fmt"
	"io"
	"log"
	"notes-back/database"
	requestTypes "notes-back/types/requestTypes"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func CreateModelRoutes[Model DefinedModel](modelInstance Model, router *gin.RouterGroup, db *mongo.Database) {
	modelType := reflect.TypeOf(modelInstance)
	modelName := modelType.Name()
	collection := database.GetCollection(db, strings.ToLower(modelName))
	modelSchema := GetModelSchema(modelInstance)
	router.GET("/"+modelName+"/:id", func(c *gin.Context) {
		getByID[Model](c, collection)
	})
	router.GET("/"+modelName, func(c *gin.Context) {
		getWithParams[Model](c, collection, modelSchema)
	})
	router.POST("/"+modelName, func(c *gin.Context) {
		post(c, collection, modelInstance)
	})
	router.PATCH("/"+modelName+"/:id", func(c *gin.Context) {
		patch[Model](c, collection)
	})
}

func getByID[Model DefinedModel](c *gin.Context, collection *mongo.Collection) {
	var response Response

	id := c.Param("id")
	item, err := database.GetItemByID(collection, id)
	if err != nil {
		log.Println(err)
		c.JSON(500, err)
		return
	}

	var result Model
	err = item.Decode(&result)

	if err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(500, response)
		return
	}
	response.Data = result

	c.JSON(200, response)
}

func getWithParams[Model DefinedModel](c *gin.Context, collection *mongo.Collection, modelSchema map[string]reflect.Kind) {
	var response Response

	var query requestTypes.GetQuery
	queryParam := c.Query("query")
	err := bson.UnmarshalExtJSON([]byte(queryParam), true, &query)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Invalid parameter 'query'")
		return
	}

	err = validateFields(modelSchema, query)
	if err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(400, response)
		return
	}

	cursor, code, err := database.GetItems(collection, query)

	if err != nil {
		response.Error = err.Error()
		c.JSON(code, response)
		return
	}

	var results []map[string]interface{}
	if err = cursor.All(context.Background(), &results); err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(500, response)
		return
	}
	response.Data = results

	c.JSON(code, results)
}

func post[Model DefinedModel](c *gin.Context, collection *mongo.Collection, modelInstance Model) {
	var response Response

	sbody, err := io.ReadAll(c.Request.Body)

	if err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(400, response)
		return
	}

	model, err := modelInstance.FromPostPayload(sbody)

	if err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(400, response)
		return
	}

	result, err := collection.InsertOne(context.Background(), model)
	if err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(500, response)
		return
	}

	response.Data = result

	c.JSON(200, response)
}

func patch[Model DefinedModel](c *gin.Context, collection *mongo.Collection) {
	id := c.Param("id")
	var body requestTypes.PatchBody
	var response Response

	if id == "" {
		response.Error = "id is required"
		c.JSON(400, response)
		return
	}

	if err := c.BindJSON(&body); err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(400, response)
		return
	}

	_, err := BytesToModel[Model]([]byte(body.Update))

	if err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(400, response)
		return
	}

	update := bson.D{}
	err = bson.UnmarshalExtJSON(body.Update, true, &update)

	if err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(400, response)
		return
	}

	result, err := database.UpdateWithID(collection, id, update, body.Operator)

	if err != nil {
		response.Error = err.Error()
		log.Println(err)
		c.JSON(500, response)
		return
	}

	response.Data = result
	c.JSON(200, response)

}

func GetModelSchema(model DefinedModel) map[string]reflect.Kind {
	modelType := reflect.TypeOf(model)
	modelSchema := make(map[string]reflect.Kind)

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		} else {
			jsonTag = strings.Split(jsonTag, ",")[0]
		}
		fieldType := field.Type.Kind()
		modelSchema[jsonTag] = fieldType
	}

	return modelSchema
}

func validateFields(modelSchema map[string]reflect.Kind, query requestTypes.GetQuery) error {

	for _, f := range query.Filters {
		_, ok := modelSchema[f.Key]
		if !ok {
			return fmt.Errorf("field '%s' not found in model schema", f.Key)
		}
	}
	for _, field := range query.Return {
		_, ok := modelSchema[field]
		if !ok {
			return fmt.Errorf("return field '%s' is not in model schema", field)
		}
	}

	return nil
}

// func validateFieldTypes(expectedType reflect.Kind, fieldValue interface{}) error {
// 	// check if fieldValue is of the expected type
// 	switch expectedType {
// 	case reflect.String:
// 		_, ok := fieldValue.(string)
// 		if !ok {
// 			return fmt.Errorf("expected string")
// 		}
// 	case reflect.Int:
// 		_, ok := fieldValue.(int)
// 		if !ok {
// 			return fmt.Errorf("expected int")
// 		}
// 	case reflect.Bool:
// 		_, ok := fieldValue.(bool)
// 		if !ok {
// 			return fmt.Errorf("expected bool")
// 		}
// 	case reflect.Float64:
// 		_, ok := fieldValue.(float64)
// 		if !ok {
// 			return fmt.Errorf("expected float64")
// 		}

// 	}
// 	return nil
// }
