package model

type Staff struct {
	Base
	Sid        string `gorm:"primary_key;index" json:"sid"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Phone      string `gorm:"default:''" json:"phone"`
}
