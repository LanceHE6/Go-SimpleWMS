package model

type Department struct {
	Base
	Did  string `gorm:"primary_key;index" json:"did"`
	Name string `gorm:"unique" json:"name"`
}
