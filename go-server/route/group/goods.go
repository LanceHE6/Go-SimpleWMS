package group

import (
	"Go_simpleWMS/handler/goods"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func GoodsGroup(ginApi *gin.RouterGroup) {
	goodsGroup := ginApi.Group("/goods", utils.AuthMiddleware())
	goodsGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goods.AddGoods(c)
	})
	goodsGroup.POST("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goods.UpdateGoods(c)
	})
	goodsGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goods.DeleteGoods(c)
	})
	goodsGroup.GET("/search", func(c *gin.Context) {
		goods.SearchGoods(c)
	})
}
