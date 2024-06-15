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
	transferGroup.GET("/search", func(c *gin.Context) {
		transfer.SearchTrans(c)
	})
	transferGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		transfer.DeleteTrans(c)
	})
	transferGroup.PUT("/audit", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		transfer.AuditTrans(c)
	})
	transferGroup.PUT("/audit/revoke", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		transfer.RevokeAudit(c)
	})
}
