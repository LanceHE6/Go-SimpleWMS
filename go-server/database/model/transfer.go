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
	OutIid        string    `gorm:"default:null" json:"out_iid"`
	InIid         string    `gorm:"default:null" json:"in_iid"`
	Operator      string    `json:"operator"`
	Comment       string    `gorm:"default:''" json:"comment"`
	Audited       bool      `gorm:"default:false" json:"audited"`
	Passed        bool      `gorm:"default:false" json:"passed"`
	Auditor       string    `gorm:"default:null" json:"auditor"`
	AuditedTime   time.Time `gorm:"default:null" json:"audited_time"`
	AuditComment  string    `gorm:"default:''" json:"audit_comment"`
}
