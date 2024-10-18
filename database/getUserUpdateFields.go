package database

import (
	"notes-back/types/requestTypes"
)

func GetUserUpdateFields(update *requestTypes.UpdateUser, fields *map[string]any) {
	if update.FirstName!= nil {
    (*fields)["first_name"] = *update.FirstName
  }
  if update.LastName!= nil {
    (*fields)["last_name"] = *update.LastName
  }
  
} 