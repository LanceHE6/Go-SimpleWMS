package model

import "time"

type MyModel struct {
	ID        uint      `gorm:"primary_key;"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

type User struct {
	MyModel
	Uid        string `gorm:"primary_key"`
	Account    string
	Password   string
	Permission int
	Nickname   string
	Phone      string `gorm:"type:varchar(100);default:''"`
	Token      string `gorm:"type:varchar(255);default:''"`
}

type Department struct {
	MyModel
	Did  string `gorm:"primary_key"`
	Name string
}
