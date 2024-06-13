package model

type GoodsType struct {
	Base
	Gtid     string `gorm:"primary_key;index" json:"gtid"`
	Name     string `json:"name"`
	TypeCode string `gorm:"default:''" json:"type_code"`
}
