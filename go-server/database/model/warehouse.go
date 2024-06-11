package model

type Warehouse struct {
	Base
	Wid     string `gorm:"primary_key;index" json:"wid"`
	Name    string `json:"name"`
	Manager string `json:"manager"`
	Comment string `gorm:"default:''" json:"comment"`
	Status  int    `gorm:"default:1" json:"status"`
}
