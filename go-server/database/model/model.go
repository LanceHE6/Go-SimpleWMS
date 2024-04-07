package model

import (
	"time"
)

type MyModel struct {
	ID        uint      `gorm:"primary_key;"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

// 数据表模型

type User struct {
	MyModel
	Uid        string `gorm:"primary_key;index"`
	Account    string
	Password   string
	Permission int
	Nickname   string
	Phone      string `gorm:"type:varchar(100);default:''"`
	Token      string `gorm:"type:varchar(255);default:''"`
}

type Department struct {
	MyModel
	Did  string `gorm:"primary_key;index"`
	Name string
}

type Unit struct {
	MyModel
	Unid string `gorm:"primary_key;index"`
	Name string
}

type Staff struct {
	MyModel
	Sid        string `gorm:"primary_key;index"`
	Name       string
	Department string
	Phone      string `gorm:"default:''"`
}

type InventoryType struct {
	MyModel
	Itid     string `gorm:"primary_key;index"`
	Name     string
	TypeCode string
}

type Warehouse struct {
	MyModel
	Wid     string `gorm:"primary_key;index"`
	Name    string
	Manager string
	Comment string `gorm:"default:''"`
	Status  int    `gorm:"default:1"`
}
