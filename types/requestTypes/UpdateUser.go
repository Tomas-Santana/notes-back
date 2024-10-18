package requestTypes

type UpdateUser struct {
	ID string `bson:"_id,omitempty" json:"_id" binding:"required"`
	FirstName *string `bson:"firstName" json:"firstName"`
	LastName *string `bson:"lastName" json:"lastName"`
}