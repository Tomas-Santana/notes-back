package database

import (
	"notes-back/types/requestTypes"
	"unicode/utf8"
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

	if _, ok := (*fields)["content"]; ok {
		var preview string
		if utf8.RuneCountInString((*fields)["content"].(string)) > 100 {
			preview = (*fields)["content"].(string)[:100]
		} else {
			preview = (*fields)["content"].(string)
		}
		(*fields)["preview"] = preview
	}
	
}