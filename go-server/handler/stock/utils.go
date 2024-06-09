package stock

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
)

func GetStock(warehouse string, goods string) float64 {
	db := myDb.GetMyDbConnection()
	var stock model.Stock
	notExist := db.Model(model.Stock{}).Where("warehouse = ? and goods = ?", warehouse, goods).First(&stock).RecordNotFound()
	if notExist {
		return 0
	}
	// 返回库存数量
	return stock.Quantity
}
