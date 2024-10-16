package database

import (
	"notes-back/types"
	"notes-back/types/requestTypes"
)

type Database interface {
	Connect()	error
	Disconnect()	error
	GetUserById(string) (types.User, error)
	GetUserByEmail(string) (types.User, error)
	CreateUser(*types.User) error
	CreateNote(string, *types.Note) (string, error)
	UpdateNote(string, *requestTypes.UpdateNote) error
	GetUserNotes(string) ([]types.Note, error)
	GetNoteById(string) (types.Note, error)
	DeleteNote([]string) error
	DeleteNoteById(string) error
	StringToId(string) (interface{}, error)
	CreateCategory(*types.Category) (error)
	GetCategories(string) ([]types.Category, error)
	DeleteCategory(string) error
}