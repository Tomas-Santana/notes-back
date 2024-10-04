package types


type User struct {
	ID        string `bson:"_id,omitempty" json:"_id"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
}