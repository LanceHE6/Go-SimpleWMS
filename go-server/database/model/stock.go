package model

type Stock struct {
	Time
	Goods     string  `json:"goods"`
	Warehouse string  `json:"warehouse"`
	Quantity  float64 `json:"quantity"`
}
