package model

type InventoryType struct {
	Base
	Itid         string `gorm:"primary_key;index" json:"itid"`
	Name         string `json:"name"`
	Type         int    `gorm:"default:1" json:"type"` // 1: 入库 2: 出库
	TypeCode     string `json:"type_code"`
	IsSystemType int    `gorm:"default:0" json:"-"` // 1: 系统类型 0: 用户自定义类型
	IsDeleted    int    `gorm:"default:0" json:"-"`
}
