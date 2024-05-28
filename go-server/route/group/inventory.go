package group

import (
	"Go_simpleWMS/handler/inventory"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func InventoryGroup(ginApi *gin.RouterGroup) {
	inventoryGroup := ginApi.Group("/inv", utils.AuthMiddleware())
	inventoryGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		inventory.AddInv(c)
	})
}
