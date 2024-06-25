package group

import (
	"Go_simpleWMS/handler/staff"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func StaffGroup(ginApi *gin.RouterGroup) {
	staffGroup := ginApi.Group("/staff", utils.AuthMiddleware())
	staffGroup.POST("/add",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("员工", "增加"),
		func(c *gin.Context) {
			staff.AddStaff(c)
		})
	staffGroup.PUT("/update",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("员工", "修改"),
		func(c *gin.Context) {
			staff.UpdateStaff(c)
		})
	staffGroup.DELETE("/delete",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("员工", "删除"),
		func(c *gin.Context) {
			staff.DeleteStaff(c)
		})
	staffGroup.GET("/list", func(c *gin.Context) {
		staff.ListStaff(c)
	})
	staffGroup.POST("/upload",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("员工", "导入"),
		func(c *gin.Context) {
			staff.UploadStaffs(c)
		})
}
