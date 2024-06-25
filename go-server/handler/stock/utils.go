package stock

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetStock(warehouse string, goods string) float64 {
	db := my_db.GetMyDbConnection()
	var stock model.Stock
	notExist := db.Model(model.Stock{}).Where("warehouse = ? and goods = ?", warehouse, goods).First(&stock).RecordNotFound()
	if notExist {
		return 0
	}
	// 返回库存数量
	return stock.Quantity
}

func UpdateStocks(GoodsList model.GoodsList, Warehouse string, inventoryType model.InventoryType, context *gin.Context, db *gorm.DB) int {
	// 更新库存表
	for _, goodsOrder := range GoodsList {
		var stock model.Stock
		// 判断库存表中是否存在这个映射
		notExist := db.Model(model.Stock{}).Where("warehouse=? AND goods=?", Warehouse, goodsOrder.Goods).First(&stock).RecordNotFound()
		var newStock = model.Stock{
			Warehouse: Warehouse,
			Goods:     goodsOrder.Goods,
			Quantity:  goodsOrder.Amount,
		}
		// 根据出入库方向执行不同的操作
		if inventoryType.Type == 1 {
			if notExist {
				// 不存在就插入记录新增映射
				db.Model(model.Stock{}).Create(&newStock)
			} else {
				// 存在就更新记录
				db.Model(model.Stock{}).Where("warehouse=? AND goods=?", Warehouse, goodsOrder.Goods).Update("quantity", stock.Quantity+goodsOrder.Amount)
			}
		} else {
			if notExist {
				// 无仓库货品映射，无法执行出库操作
				context.JSON(http.StatusBadRequest, response.Response(405, "No warehouse-goods mapping", nil))
				return 1
			} else {
				// 存在就更新记录
				if stock.Quantity < goodsOrder.Amount {
					context.JSON(http.StatusBadRequest, response.Response(406, "Insufficient inventory", nil))
					return 1
				} else {
					var updateData = map[string]interface{}{
						"quantity": stock.Quantity - goodsOrder.Amount,
					}
					db.Model(model.Stock{}).Where("warehouse=? AND goods=?", Warehouse, goodsOrder.Goods).Updates(updateData)
				}
			}
		}
	}
	return 0
}

func GetGoodsListAmountByGoods(goodsList model.GoodsList, goods string) float64 {
	var amount float64 = 0
	for _, goodsOrder := range goodsList {
		if goodsOrder.Goods == goods {
			amount += goodsOrder.Amount
		}
	}
	return amount
}
