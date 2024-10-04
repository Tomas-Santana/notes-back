package models
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID	primitive.ObjectID	`bson:"_id,omitempty" json:"_id"`
	Email	string	`bson:"email" json:"email"`
	Password	string	`bson:"password" json:"password"`
	FirstName	string	`bson:"firstName" json:"firstName"`
	LastName	string	`bson:"lastName" json:"lastName"`
}