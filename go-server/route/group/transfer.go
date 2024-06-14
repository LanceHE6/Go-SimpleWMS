package group

import (
	"Go_simpleWMS/handler/transfer"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func TransferGroup(ginApi *gin.RouterGroup) {
	transferGroup := ginApi.Group("/trans", utils.AuthMiddleware())
	transferGroup.POST("/add", utils.IsAdminMiddleware(), func(c *gin.Context) {
		transfer.AddTrans(c)
	})
	//transferGroup.GET("/search", func(c *gin.Context) {
	//	transfer.SearchInv(c)
	//})
	//transferGroup.DELETE("/delete", utils.IsAdminMiddleware(), func(c *gin.Context) {
	//	transfer.DeleteInv(c)
	//})
}
