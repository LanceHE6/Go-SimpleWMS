package group

import (
	"Go_simpleWMS/handler/opLog"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func OPLogGroup(ginApi *gin.RouterGroup) {
	opLogGroup := ginApi.Group("/log", utils.AuthMiddleware())
	opLogGroup.DELETE("/clean",
		utils.IsSuperAdminMiddleware(),
		func(c *gin.Context) {
			opLog.CleanOPLog(c)
		})
	opLogGroup.GET("/search", func(c *gin.Context) {
		opLog.SearchOPLog(c)
	})

}
