package group

import (
	"Go_simpleWMS/handler/goods_type"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func GoodsType(ginApi *gin.RouterGroup) {
	goodsTypeGroup := ginApi.Group("/gt", utils.AuthMiddleware())
	goodsTypeGroup.POST("/add",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("货品类型", "增加"),
		func(c *gin.Context) {
			goods_type.AddGoodsType(c)
		})
	goodsTypeGroup.PUT("/update",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("货品类型", "修改"),
		func(c *gin.Context) {
			goods_type.UpdateGoodsType(c)
		})
	goodsTypeGroup.DELETE("/delete",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("货品类型", "删除"),
		func(c *gin.Context) {
			goods_type.DeleteGoodsType(c)
		})
	goodsTypeGroup.GET("/list", func(c *gin.Context) {
		goods_type.ListGoodsType(c)
	})
}
