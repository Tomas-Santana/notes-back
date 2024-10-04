package types

import (
	"time"
)

type Note struct {
	ID         string      `bson:"_id,omitempty" json:"_id"`
	Title      string      `bson:"title" json:"title"`
	Preview    string      `bson:"preview" json:"preview"`
	Content    string      `bson:"content" json:"content"`
	Html       string      `bson:"html" json:"html"`
	IsFavorite bool        `bson:"isFavorite" json:"isFavorite"`
	CreatedAt  time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time   `bson:"updatedAt" json:"updatedAt"`
	UserID     interface{} `bson:"userID" json:"userID"`
}
