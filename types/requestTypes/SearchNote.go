package requestTypes

type SearchNotes struct {
	Query string `json:"query" validate:"required,min=3,max=100"`
}
