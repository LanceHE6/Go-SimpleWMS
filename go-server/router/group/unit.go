package group

import (
	"Go_simpleWMS/handler/unit"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func UnitGroup(ginApi *gin.RouterGroup) {
	unitGroup := ginApi.Group("/unit", utils.AuthMiddleware())
	unitGroup.POST("/add",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("计量单位", "增加"),
		func(c *gin.Context) {
			unit.AddUnit(c)
		})
	unitGroup.DELETE("/delete",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("计量单位", "删除"),
		func(c *gin.Context) {
			unit.DeleteUnit(c)
		})
	unitGroup.GET("/list", func(c *gin.Context) {
		unit.ListUnit(c)
	})
	unitGroup.PUT("/update",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("计量单位", "修改"),
		func(c *gin.Context) {
			unit.UpdateUnit(c)
		})
}
