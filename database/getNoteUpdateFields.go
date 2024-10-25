package database

import (
	"notes-back/types/requestTypes"
)

func GetNoteUpdateFields(update *requestTypes.UpdateNote, fields *map[string]any) {

	if update.Title != nil {
		(*fields)["title"] = *update.Title
	}

	if update.Content != nil {
		(*fields)["content"] = *update.Content
	}

	if update.Html != nil {
		(*fields)["html"] = *update.Html
	}

	if update.IsFavorite != nil {
		(*fields)["isFavorite"] = *update.IsFavorite
	}

	if update.Importance != nil {
		(*fields)["importance"] = *update.Importance
	}

	if update.Categories != nil {
		(*fields)["categories"] = *update.Categories
	}

	if update.UpdatedAt != nil {
		(*fields)["updatedAt"] = *update.UpdatedAt
	}
	
}