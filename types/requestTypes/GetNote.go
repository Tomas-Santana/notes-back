package requestTypes

type GetNote struct {
	ID string `json:"id" validation:"required"`
}