package group

import (
	"Go_simpleWMS/handler/transfer"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func TransferGroup(ginApi *gin.RouterGroup) {
	transferGroup := ginApi.Group("/trans", utils.AuthMiddleware())
	transferGroup.POST("/add",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("调拨单", "增加"),
		func(c *gin.Context) {
			transfer.AddTrans(c)
		})
	transferGroup.GET("/search", func(c *gin.Context) {
		transfer.SearchTrans(c)
	})
	transferGroup.DELETE("/delete",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("调拨单", "删除"),
		func(c *gin.Context) {
			transfer.DeleteTrans(c)
		})
	transferGroup.PUT("/audit",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("调拨单", "审核"),
		func(c *gin.Context) {
			transfer.AuditTrans(c)
		})
	transferGroup.PUT("/audit/revoke",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("调拨单", "撤销审核"),
		func(c *gin.Context) {
			transfer.RevokeAudit(c)
		})
}
