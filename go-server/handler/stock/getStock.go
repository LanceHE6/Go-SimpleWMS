package stock

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStockRequest(context *gin.Context) {
	// 获取库存信息
	goods := context.Query("goods")
	warehouse := context.Query("warehouse")
	if warehouse == "" && goods == "" {
		// 获取所有的仓库的所有物品库存信息
		db := myDb.GetMyDbConnection()
		type goodsSummary struct {
			Goods    string  `json:"goods"`
			Quantity float64 `json:"quantity"`
		}
		var result []goodsSummary
		db.Table("stocks").
			Select("goods, SUM(quantity) AS quantity").
			Group("goods").
			Find(&result)
		context.JSON(http.StatusOK, response.Response(201, "Get stock success", gin.H{
			"rows":  result,
			"total": len(result),
		}))
		return
	}
	// 获取仓库的库存信息
	if goods == "" {
		var stocks []model.Stock
		db := myDb.GetMyDbConnection()
		err := db.Model(model.Stock{}).Where("warehouse = ?", warehouse).Find(&stocks).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Get stock failed", err.Error()))
			return
		}

		context.JSON(http.StatusOK, response.Response(202, "Get stock success", gin.H{
			"rows":  stocks,
			"total": len(stocks),
		}))
		return
	}
	// 获取商品的总库存信息
	if warehouse == "" {
		var stocks []model.Stock
		db := myDb.GetMyDbConnection()
		err := db.Model(model.Stock{}).Where("goods = ?", goods).Find(&stocks).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(502, "Get stock failed", err.Error()))
			return
		}
		// 计算总库存
		var totalQuantity float64
		for _, stock := range stocks {
			totalQuantity += stock.Quantity
		}

		context.JSON(http.StatusOK, response.Response(203, "Get stock success", gin.H{
			"total_quantity": totalQuantity,
			"rows":           stocks,
		}))
		return
	}
	// 获取指定仓库的指定商品的库存信息
	quantity := GetStock(warehouse, goods)

	context.JSON(http.StatusOK, response.Response(204, "Get stock success", gin.H{
		"goods":     goods,
		"warehouse": warehouse,
		"quantity":  quantity,
	}))
}
