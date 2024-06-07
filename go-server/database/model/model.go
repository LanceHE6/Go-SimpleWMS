package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Time struct {
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

type MyModel struct {
	ID uint `gorm:"primary_key;" json:"id"`
	Time
}

// 数据表模型

type User struct {
	MyModel
	Uid        string `gorm:"primary_key;index" json:"uid"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	Permission int    `json:"permission"`
	Nickname   string `json:"nickname"`
	Phone      string `gorm:"type:varchar(100);default:''" json:"phone"`
	Token      string `gorm:"type:varchar(255);default:''" json:"token"`
}

type Department struct {
	MyModel
	Did  string `gorm:"primary_key;index" json:"did"`
	Name string `gorm:"unique" json:"name"`
}

type Unit struct {
	MyModel
	Unid string `gorm:"primary_key;index" json:"unid"`
	Name string `json:"name"`
}

type Staff struct {
	MyModel
	Sid        string `gorm:"primary_key;index" json:"sid"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Phone      string `gorm:"default:''" json:"phone"`
}

type InventoryType struct {
	MyModel
	Itid     string `gorm:"primary_key;index" json:"itid"`
	Name     string `json:"name"`
	Type     int    `gorm:"default:1" json:"type"` // 1: 入库 2: 出库
	TypeCode string `json:"type_code"`
}

type Warehouse struct {
	MyModel
	Wid     string `gorm:"primary_key;index" json:"wid"`
	Name    string `json:"name"`
	Manager string `json:"manager"`
	Comment string `gorm:"default:''" json:"comment"`
	Status  int    `gorm:"default:1" json:"status"`
}

type GoodsType struct {
	MyModel
	Gtid     string `gorm:"primary_key;index" json:"gtid"`
	Name     string `json:"name"`
	TypeCode string `gorm:"default:''" json:"type_code"`
}

type Goods struct {
	MyModel
	Gid          string  `gorm:"primary_key;index" json:"gid"`
	GoodsCode    string  `gorm:"unique" json:"goods_code"`
	Name         string  `json:"name"`
	Model        string  `json:"model"`
	GoodsType    string  `gorm:"default:''" json:"goods_type"`
	Manufacturer string  `gorm:"default:''" json:"manufacturer"`
	Unit         string  `json:"unit"`
	Image        string  `gorm:"default:''" json:"image"`
	UnitPrice    float64 `gorm:"default:0" json:"unit_price"`
}

// GoodsOrder 非建表数据结构
type GoodsOrder struct {
	Goods   string  `json:"goods"`
	Amount  float64 `json:"amount"`
	Comment string  `json:"comment"`
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
	Iid           string    `gorm:"primary_key;index" json:"iid"`
	Number        string    `gorm:"unique" json:"number"`
	Date          time.Time `json:"date"`
	GoodsList     GoodsList `gorm:"type:json" json:"goods_list"`
	OldGoodsList  GoodsList `gorm:"type:json" json:"old_goods_list"` // 更新前的库存信息
	NewGoodsList  GoodsList `gorm:"type:json" json:"new_goods_list"` // 更新后的库存信息
	Warehouse     string    `json:"warehouse"`
	InventoryType string    `json:"inventory_type"`
	Department    string    `gorm:"default:''" json:"department"`
	Operator      string    `json:"operator"`
	Comment       string    `gorm:"default:''" json:"comment"`
	Manufacturer  string    `gorm:"default:''" json:"manufacturer"`
}

type Stock struct {
	Time
	Goods     string  `json:"goods"`
	Warehouse string  `json:"warehouse"`
	Quantity  float64 `json:"quantity"`
}
