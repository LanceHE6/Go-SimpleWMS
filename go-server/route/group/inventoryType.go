package group

import (
	"Go_simpleWMS/handler/inventoryType"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func InventoryTypeGroup(ginApi *gin.RouterGroup) {
	inventoryTypeGroup := ginApi.Group("/invt", utils.AuthMiddleware())
	inventoryTypeGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		inventoryType.AddInventoryType(c)
	})
	inventoryTypeGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		inventoryType.UpdateInventoryType(c)
	})
	inventoryTypeGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		inventoryType.DeleteInventoryType(c)
	})
	inventoryTypeGroup.GET("/list", func(c *gin.Context) {
		inventoryType.ListInventoryType(c)
	})
}
