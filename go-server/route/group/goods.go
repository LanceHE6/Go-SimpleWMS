package group

import (
	"Go_simpleWMS/handler/goods"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func GoodsGroup(ginApi *gin.RouterGroup) {
	goodsGroup := ginApi.Group("/goods", utils.AuthMiddleware())
	goodsGroup.POST("/add",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("货品", "添加"),
		func(c *gin.Context) {
			goods.AddGoods(c)
		})
	goodsGroup.PUT("/update",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("货品", "修改"),
		func(c *gin.Context) {
			goods.UpdateGoods(c)
		})
	goodsGroup.DELETE("/delete",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("货品", "删除"),
		func(c *gin.Context) {
			goods.DeleteGoods(c)
		})
	goodsGroup.GET("/search", func(c *gin.Context) {
		goods.SearchGoods(c)
	})
	goodsGroup.DELETE("/delete/image",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("货品", "删除货品图片"),
		func(c *gin.Context) {
			goods.DeleteGoodsAttachment(c, goods.IMAGE)
		})
	goodsGroup.DELETE("/delete/file",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("货品", "删除货品附件"),
		func(c *gin.Context) {
			goods.DeleteGoodsAttachment(c, goods.FILE)
		})
}
