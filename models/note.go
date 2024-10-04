package models

import (
	// mongo primitive package
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	Title  string `bson:"title" json:"title"`
	Preview string `bson:"preview" json:"preview"`
	Content string `bson:"content" json:"content"`
	IsFavorite bool `bson:"isFavorite" json:"isFavorite"`
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
}
func (Note) FromPostPayload(payload []byte) (DefinedModel, error) {

	note, err := BytesToModel[Note](payload)

	if err != nil {
		return Note{}, err
	}
	if note.Title == "" {
		return Note{}, errors.New("title is required")
	}
	if note.Preview == "" {
		return Note{}, errors.New("preview is required")
	}
	if note.Content == "" {
		return Note{}, errors.New("content is required")
	}
	
	if note.CreatedAt == primitive.DateTime(0) {
		note.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	}
	
	return note, nil

}





