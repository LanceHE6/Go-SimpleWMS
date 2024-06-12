package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
)

type File struct {
	Path string `json:"path"`
}
type FileList []File

func (fl *FileList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to cast value to []byte")
	}

	err := json.Unmarshal(bytes, fl)
	if err != nil {
		return err
	}

	return nil
}

func (fl FileList) Value() (driver.Value, error) {
	bytes, err := json.Marshal(fl)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

type Goods struct {
	Base
	Gid          string   `gorm:"primary_key;index" json:"gid"`
	GoodsCode    string   `gorm:"unique" json:"goods_code"`
	Name         string   `json:"name"`
	Model        string   `json:"model"`
	GoodsType    string   `gorm:"default:''" json:"goods_type"`
	Manufacturer string   `gorm:"default:''" json:"manufacturer"`
	Unit         string   `json:"unit"`
	Images       FileList `gorm:"type:json" json:"images"`
	Files        FileList `gorm:"type:json" json:"files"`
	UnitPrice    float64  `gorm:"default:0" json:"unit_price"`
}

// BeforeDelete 定义钩子，在删除商品之前，删除商品的库存信息以及出入库单中包含该货品的记录
func (g *Goods) BeforeDelete(tx *gorm.DB) (err error) {
	var invs []Inventory
	tx.Model(Inventory{}).Where("json_extract(goods_list, '$[*].goods') LIKE ?", "%"+g.Gid+"%").Find(&invs)
	for _, inv := range invs {
		// 如果goodsList里面只有含有该货品的记录则删除整个出入库单
		if len(inv.GoodsList) == 1 {
			err = tx.Delete(&inv).Error
			continue
		}
		var goodsList GoodsList
		for _, goodsOrder := range inv.GoodsList {
			if goodsOrder.Goods != g.Gid {
				goodsList = append(goodsList, goodsOrder)
			}
		}
		var oldGoodsList GoodsList
		for _, goodsOrder := range inv.OldGoodsList {
			if goodsOrder.Goods != g.Gid {
				oldGoodsList = append(oldGoodsList, goodsOrder)
			}
		}
		var newGoodsList GoodsList
		for _, goodsOrder := range inv.NewGoodsList {
			if goodsOrder.Goods != g.Gid {
				newGoodsList = append(newGoodsList, goodsOrder)
			}
		}
		inv.GoodsList = goodsList
		inv.OldGoodsList = oldGoodsList
		inv.NewGoodsList = newGoodsList

		// 更新出入库表
		err = tx.Save(&inv).Error
	}
	// 删除库存
	return tx.Delete(&Stock{}, "goods = ?", g.Gid).Error
}
