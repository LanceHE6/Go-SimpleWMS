package group

import (
	"Go_simpleWMS/handler/user"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func UserGroup(ginApi *gin.RouterGroup) {
	userGroup := ginApi.Group("/user")

	userGroup.POST("/register", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		user.Register(c)
	})
	userGroup.POST("/upload", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		user.UploadUsers(c)
	})
	userGroup.POST("/login", func(c *gin.Context) {
		user.Login(c)
	})
	userGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		user.DeleteUser(c)
	})
	userGroup.PUT("/update", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		user.UpdateUser(c)
	})
	userGroup.GET("/list", utils.AuthMiddleware(), func(c *gin.Context) {
		user.ListUsers(c)
	})

	emailGroup := userGroup.Group("/email")

	emailGroup.POST("/bind", utils.AuthMiddleware(), func(c *gin.Context) {
		user.BindEmail(c)
	})
	emailGroup.POST("/verify", utils.AuthMiddleware(), func(c *gin.Context) {
		user.VerifyEmail(c)
	})
}
