package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Time struct {
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
type MyModel struct {
	ID uint `gorm:"primary_key;"`
	Time
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
	Manufacturer string `gorm:"default:''"`
	Unit         string
	Image        string  `gorm:"default:''"`
	UnitPrice    float64 `gorm:"default:0"`
}

// GoodsOrder 非建表数据结构
type GoodsOrder struct {
	Goods   string
	Amount  float64
	Comment string
}

// GoodsList 定义自定义类型 GoodsList
type GoodsList []GoodsOrder

func (gol *GoodsList) Scan(value interface{}) error {
	// 假设value是一个[]byte（从数据库读取的JSON数据）
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to cast value to []byte")
	}

	// 将JSON数据解码到gol指向的GoodsList实例中
	err := json.Unmarshal(bytes, gol)
	if err != nil {
		return err
	}

	return nil
}

func (gol GoodsList) Value() (driver.Value, error) {
	// 将GoodsList编码为JSON格式的[]byte
	bytes, err := json.Marshal(gol)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

type Inventory struct {
	MyModel
	Iid           string `gorm:"primary_key;index"`
	Number        string `gorm:"unique"`
	Date          time.Time
	GoodsList     GoodsList `gorm:"type:json"`
	OldGoodsList  GoodsList `gorm:"type:json"` // 更新前的库存信息
	NewGoodsList  GoodsList `gorm:"type:json"` // 更新后的库存信息
	Warehouse     string
	InventoryType string
	Department    string `gorm:"default:''"`
	Operator      string
	Comment       string `gorm:"default:''"`
	Manufacturer  string `gorm:"default:''"`
}

type Stock struct {
	Time
	Goods     string
	Warehouse string
	Quantity  float64
}
