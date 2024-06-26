package group

import (
	"Go_simpleWMS/handler/warehouse"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func WarehouseGroup(ginApi *gin.RouterGroup) {
	warehouseGroup := ginApi.Group("/warehouse", utils.AuthMiddleware())
	warehouseGroup.POST("/add",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("仓库", "增加"),
		func(c *gin.Context) {
			warehouse.AddWarehouse(c)
		})
	warehouseGroup.DELETE("/delete",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("仓库", "删除"),
		func(c *gin.Context) {
			warehouse.DeleteWarehouse(c)
		})
	warehouseGroup.PUT("/update",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("仓库", "修改"),
		func(c *gin.Context) {
			warehouse.UpdateWarehouse(c)
		})
	warehouseGroup.GET("/list", func(c *gin.Context) {
		warehouse.ListWarehouse(c)
	})
}
