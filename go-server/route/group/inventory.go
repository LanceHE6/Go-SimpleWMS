package group

import (
	"Go_simpleWMS/handler/inventory"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func InventoryGroup(ginApi *gin.RouterGroup) {
	inventoryGroup := ginApi.Group("/inv", utils.AuthMiddleware())
	inventoryGroup.POST("/add", utils.IsAdminMiddleware(), func(c *gin.Context) {
		inventory.AddInv(c)
	})
	inventoryGroup.GET("/search", func(c *gin.Context) {
		inventory.SearchInv(c)
	})
}
