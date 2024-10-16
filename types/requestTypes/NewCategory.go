package requestTypes

type NewCategory struct {
	Name string `json:"name" validate:"required"`
}