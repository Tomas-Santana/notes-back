package database

import (
	"notes-back/types/requestTypes"
)

func GetUserUpdateFields(update *requestTypes.UpdateUser, fields *map[string]any) {
	if update.FirstName!= nil {
    (*fields)["firstName"] = *update.FirstName
  }
  if update.LastName!= nil {
    (*fields)["lastName"] = *update.LastName
  }
  
} 