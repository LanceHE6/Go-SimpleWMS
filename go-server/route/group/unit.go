package group

import (
	"Go_simpleWMS/handler/unit"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func UnitGroup(ginApi *gin.RouterGroup) {
	unitGroup := ginApi.Group("/unit", utils.AuthMiddleware())
	unitGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		unit.AddUnit(c)
	})
	unitGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		unit.DeleteUnit(c)
	})
	unitGroup.GET("/list", func(c *gin.Context) {
		unit.ListUnit(c)
	})
}
