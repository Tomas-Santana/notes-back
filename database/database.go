package database

import "notes-back/types"

type Database interface {
	Connect()	error
	Disconnect()	error
	GetUserById(string) (types.User, error)
	GetUserByEmail(string) (types.User, error)
	CreateUser(*types.User) error
	CreateNote(string, *types.Note) (string, error)
	GetUserNotes(string) ([]types.Note, error)
	GetNoteById(string) (types.Note, error)
	StringToId(string) (interface{}, error)
}