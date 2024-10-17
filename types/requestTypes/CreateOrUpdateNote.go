package requestTypes

import "notes-back/types"

type CreateNote struct {
	ID         string `bson:"_id,omitempty" json:"_id"`
	Title      string `json:"title" binding:"required,min=1,max=100"`
	Content    string `json:"content" binding:"required,min=1,max=1000"`
	Html       string `json:"html" binding:"required,min=3,max=1000"`
	Importance int    `json:"importance" binding:"min=0,max=5"`
	Categories []types.NoteCategory `json:"categories"`
}

type UpdateNote struct {
	ID         string  `bson:"_id,omitempty" json:"_id" binding:"required"`
	Title      *string `json:"title" binding:"omitempty,min=1,max=100"`
	Content    *string `json:"content" binding:"omitempty,min=1,max=1000"`
	Html       *string `json:"html" binding:"omitempty,min=3,max=1000"`
	IsFavorite *bool   `json:"isFavorite"`
	Importance *int    `json:"importance" binding:"min=0,max=5"`
	Categories *[]types.NoteCategory `json:"categories"`
}

type DeleteNote struct {
	ID string `bson:"_id,omitempty" json:"_id" binding:"required"`
}
