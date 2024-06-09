package stock

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStockRequest(context *gin.Context) {
	// 获取库存信息
	goods := context.Query("goods")
	warehouse := context.Query("warehouse")
	if warehouse == "" && goods == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters",
			"code":    400,
		})
		return
	}
	if goods == "" {
		var stocks []model.Stock
		db := myDb.GetMyDbConnection()
		err := db.Model(model.Stock{}).Where("warehouse = ?", warehouse).Find(&stocks).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "Get stock failed",
				"code":    500,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Get stock success",
			"code":    202,
			"rows":    stocks,
		})
		return
	}
	quantity := GetStock(warehouse, goods)

	context.JSON(http.StatusOK, gin.H{
		"message": "Get stock success",
		"code":    201,
		"data": gin.H{
			"goods":     goods,
			"warehouse": warehouse,
			"quantity":  quantity,
		},
	})
}
