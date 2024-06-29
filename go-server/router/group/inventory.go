package group

import (
	"Go_simpleWMS/handler/inventory"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func InventoryGroup(ginApi *gin.RouterGroup) {
	inventoryGroup := ginApi.Group("/inv", utils.AuthMiddleware())
	inventoryGroup.POST("/add",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("出入库单", "增加"),
		func(c *gin.Context) {
			inventory.AddInv(c)
		})
	inventoryGroup.GET("/search", func(c *gin.Context) {
		inventory.SearchInv(c)
	})
	inventoryGroup.PUT("/update",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("出入库单", "修改"),
		func(c *gin.Context) {
			inventory.UpdateInv(c)
		})
	inventoryGroup.DELETE("/delete",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("出入库单", "删除"),
		func(c *gin.Context) {
			inventory.DeleteInv(c)
		})
}
