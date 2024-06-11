package model

type Unit struct {
	Base
	Unid string `gorm:"primary_key;index" json:"unid"`
	Name string `json:"name"`
}
