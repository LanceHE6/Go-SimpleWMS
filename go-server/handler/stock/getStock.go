package stock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStockRequest(context *gin.Context) {
	// 获取库存信息
	goods := context.Query("goods")
	warehouse := context.Query("warehouse")
	quantity := GetStock(warehouse, goods)

	context.JSON(http.StatusOK, gin.H{
		"message": "Get stock success",
		"code":    200,
		"data": gin.H{
			"goods":     goods,
			"warehouse": warehouse,
			"quantity":  quantity,
		},
	})
}
