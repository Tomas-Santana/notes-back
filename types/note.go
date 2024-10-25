package types

import (
	"time"
)

type Note struct {
	ID         string      `bson:"_id,omitempty" json:"_id"`
	Title      string      `bson:"title" json:"title"`
	Content    string      `bson:"content" json:"content"`
	Html       string      `bson:"html" json:"html"`
	IsFavorite bool        `bson:"isFavorite" json:"isFavorite"`
	Importance int         `bson:"importance" json:"importance"`
	CreatedAt  time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time   `bson:"updatedAt" json:"updatedAt"`
	UserID     interface{} `bson:"userID" json:"userID"`
	Categories   []NoteCategory `bson:"categories" json:"categories"`
}
