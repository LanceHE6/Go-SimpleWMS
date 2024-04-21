package group

import (
	"Go_simpleWMS/handler/department"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func DepartmentGroup(ginApi *gin.RouterGroup) {
	departmentGroup := ginApi.Group("/dept", utils.AuthMiddleware())
	departmentGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		department.AddDepartment(c)
	})
	departmentGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		department.UpdateDepartment(c)
	})
	departmentGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		department.DeleteDepartment(c)
	})
	departmentGroup.GET("/list", func(c *gin.Context) {
		department.ListDepartment(c)
	})
}
