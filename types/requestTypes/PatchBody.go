package requestTypes

import "encoding/json"

type PatchBody struct {
	Operator string `json:"operator"`
	Update   json.RawMessage `json:"update"`
}