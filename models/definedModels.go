package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// type that is either a Note or a User
type DefinedModel interface {
	FromPostPayload(content []byte) (DefinedModel, error)
}

func BytesToModel[T DefinedModel](content []byte) (T, error) {

	var model T

	err := bson.UnmarshalExtJSON(content, false, &model)

	if err != nil {
		return model, err
	}

	fmt.Println(model)

	return model, nil
	

}
	