package group

import (
	"Go_simpleWMS/handler/goodsType"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func GoodsType(ginApi *gin.RouterGroup) {
	goodsTypeGroup := ginApi.Group("/gt", utils.AuthMiddleware())
	goodsTypeGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goodsType.AddGoodsType(c)
	})
	goodsTypeGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goodsType.UpdateGoodsType(c)
	})
	goodsTypeGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goodsType.DeleteGoodsType(c)
	})
	goodsTypeGroup.GET("/list", func(c *gin.Context) {
		goodsType.ListGoodsType(c)
	})
}
