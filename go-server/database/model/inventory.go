package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

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
	Base
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
