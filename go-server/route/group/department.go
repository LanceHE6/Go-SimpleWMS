package group

import (
	"Go_simpleWMS/handler/department"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func DepartmentGroup(ginApi *gin.RouterGroup) {
	departmentGroup := ginApi.Group("/dept", utils.AuthMiddleware())
	departmentGroup.POST("/add",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("部门", "添加"),
		func(c *gin.Context) {
			department.AddDepartment(c)
		})
	departmentGroup.PUT("/update",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("部门", "更新"),
		func(c *gin.Context) {
			department.UpdateDepartment(c)
		})
	departmentGroup.DELETE("/delete",
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("部门", "删除"),
		func(c *gin.Context) {
			department.DeleteDepartment(c)
		})
	departmentGroup.GET("/list", func(c *gin.Context) {
		department.ListDepartment(c)
	})
}
