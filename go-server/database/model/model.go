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
	Name string `gorm:"unique"`
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
	Type     int `gorm:"default:1"` // 1: 入库 2: 出库
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

type GoodsType struct {
	MyModel
	Gtid     string `gorm:"primary_key;index"`
	Name     string
	TypeCode string `gorm:"default:''"`
}

type Goods struct {
	MyModel
	Gid          string `gorm:"primary_key;index"`
	GoodsCode    string `gorm:"unique"`
	Name         string
	Model        string
	GoodsType    string `gorm:"default:''"`
	Warehouse    string
	Manufacturer string `gorm:"default:''"`
	Unit         string
	Image        string `gorm:"default:''"`
	Quantity     int    `gorm:"default:0"`
}

type Inventory struct {
	MyModel
	Iid           string `gorm:"primary_key;index"`
	Number        string `gorm:"unique"`
	Goods         string
	Amount        int
	InventoryType string
	Warehouse     string
	Operator      string
	Comment       string `gorm:"default:''"`
	Manufacturer  string `gorm:"default:''"`
}
