package requestTypes


type QueryFilter struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
	Comparison string `json:"comparison"`
	ConvertToDateTime bool `json:"convertToDateTime"`
}
type GetQuery struct {
	Filters []QueryFilter `json:"filters"`
	LogicalOperator string `json:"logicalFilter"`
	Return []string `json:"return"` 
}