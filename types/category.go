package types

type Category struct {
	ID     string      `bson:"_id,omitempty" json:"_id"`
	Name   string      `bson:"name" json:"name"`
	UserID interface{} `bson:"userID" json:"userID"`
}

type NoteCategory struct {
	ID    string `bson:"_id,omitempty" json:"_id"`
	Name  string `bson:"name" json:"name"`
}
