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
		context.JSON(http.StatusBadRequest, response.Response(401, "Missing parameters", nil))
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
			"rows": stocks,
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
			"total": totalQuantity,
			"rows":  stocks,
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
