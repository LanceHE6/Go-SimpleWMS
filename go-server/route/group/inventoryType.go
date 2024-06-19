package group

import (
	"Go_simpleWMS/handler/inventoryType"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func InventoryTypeGroup(ginApi *gin.RouterGroup) {
	inventoryTypeGroup := ginApi.Group("/invt", utils.AuthMiddleware())
	inventoryTypeGroup.POST("/add",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("出入库类型", "增加"),
		func(c *gin.Context) {
			inventoryType.AddInventoryType(c)
		})
	inventoryTypeGroup.PUT("/update",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("出入库类型", "修改"),
		func(c *gin.Context) {
			inventoryType.UpdateInventoryType(c)
		})
	inventoryTypeGroup.DELETE("/delete",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("出入库类型", "删除"),
		func(c *gin.Context) {
			inventoryType.DeleteInventoryType(c)
		})
	inventoryTypeGroup.GET("/list", func(c *gin.Context) {
		inventoryType.ListInventoryType(c)
	})
}
