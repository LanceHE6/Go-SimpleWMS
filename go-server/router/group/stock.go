package group

import (
	"Go_simpleWMS/handler/stock"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func StockGroup(ginApi *gin.RouterGroup) {
	staffGroup := ginApi.Group("/stock", utils.AuthMiddleware())
	staffGroup.GET("/get", func(c *gin.Context) {
		stock.GetStockRequest(c)
	})
}
