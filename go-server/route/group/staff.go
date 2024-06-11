package group

import (
	"Go_simpleWMS/handler/staff"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func StaffGroup(ginApi *gin.RouterGroup) {
	staffGroup := ginApi.Group("/staff", utils.AuthMiddleware())
	staffGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		staff.AddStaff(c)
	})
	staffGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		staff.UpdateStaff(c)
	})
	staffGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		staff.DeleteStaff(c)
	})
	staffGroup.GET("/list", func(c *gin.Context) {
		staff.ListStaff(c)
	})
	staffGroup.POST("/upload", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		staff.UploadStaffs(c)
	})
}
