package model

import "time"

type Transfer struct {
	Base
	Tid           string    `gorm:"primary_key;index" json:"tid"`
	Number        string    `gorm:"unique" json:"number"`
	Date          time.Time `json:"date"`
	GoodsList     GoodsList `gorm:"type:json" json:"goods_list"`
	SourWarehouse string    `json:"source_warehouse"`
	DestWarehouse string    `json:"destination_warehouse"`
	OutIid        string    `json:"out_iid"`
	InIid         string    `json:"in_iid"`
	Operator      string    `json:"operator"`
	Comment       string    `gorm:"default:''" json:"comment"`
	Checked       bool      `gorm:"default:false" json:"checked"`
	Checker       string    `json:"checker"`
	CheckTime     time.Time `gorm:"default:null" json:"check_time"`
	CheckComment  string    `gorm:"default:''" json:"check_comment"`
}
