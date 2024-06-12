package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Warehouse struct {
	Base
	Wid     string `gorm:"primary_key;index" json:"wid"`
	Name    string `json:"name"`
	Manager string `json:"manager"`
	Comment string `gorm:"default:''" json:"comment"`
	Status  int    `gorm:"default:1" json:"status"`
}

// BeforeDelete 定义钩子，在删除仓库后删除跟该仓库有关的所有库存记录以及出入库单据
func (w *Warehouse) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("BeforeDelete Hook Called for Warehouse:", w.Wid)
	tx.Delete(&Inventory{}, "warehouse = ?", w.Wid)
	return tx.Delete(&Stock{}, "warehouse = ?", w.Wid).Error
}
