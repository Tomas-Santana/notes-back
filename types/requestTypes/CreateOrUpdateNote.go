package requestTypes

type CreateNote struct {
	ID      string `bson:"_id,omitempty" json:"_id"`
	Title   string `json:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required,min=3,max=1000"`
	Html    string `json:"html" binding:"required,min=3,max=1000"`
}

type UpdateNote struct {
	ID      string `bson:"_id,omitempty" json:"_id" binding:"required"`
	Title   *string `json:"title" binding:"omitempty,min=3,max=100"`
	Content *string `json:"content" binding:"omitempty,min=3,max=1000"`
	Html    *string `json:"html" binding:"omitempty,min=3,max=1000"`
	IsFavorite *bool `json:"isFavorite" binding:"omitempty"`
}
