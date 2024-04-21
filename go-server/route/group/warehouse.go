package group

import (
	"Go_simpleWMS/handler/warehouse"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func WarehouseGroup(ginApi *gin.RouterGroup) {
	warehouseGroup := ginApi.Group("/warehouse", utils.AuthMiddleware())
	warehouseGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		warehouse.AddWarehouse(c)
	})
	warehouseGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		warehouse.DeleteWarehouse(c)
	})
	warehouseGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		warehouse.UpdateWarehouse(c)
	})
	warehouseGroup.GET("/list", func(c *gin.Context) {
		warehouse.ListWarehouse(c)
	})
}
