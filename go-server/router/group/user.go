package group

import (
	"Go_simpleWMS/handler/user"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func UserGroup(ginApi *gin.RouterGroup) {
	userGroup := ginApi.Group("/user")

	userGroup.POST("/register",
		utils.AuthMiddleware(),
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("用户", "注册"),
		func(c *gin.Context) {
			user.Register(c)
		})
	userGroup.POST("/upload",
		utils.AuthMiddleware(),
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("用户", "导入"),
		func(c *gin.Context) {
			user.UploadUsers(c)
		})
	userGroup.POST("/login",
		func(c *gin.Context) {
			user.Login(c)
		})
	userGroup.DELETE("/delete",
		utils.AuthMiddleware(),
		utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("用户", "删除"),
		func(c *gin.Context) {
			user.DeleteUser(c)
		})
	userGroup.PUT("/update",
		utils.AuthMiddleware(),
		//utils.IsSuperAdminMiddleware(),
		utils.OPLoggerMiddleware("用户", "更新"),
		func(c *gin.Context) {
			user.UpdateUser(c)
		})
	userGroup.GET("/list", utils.AuthMiddleware(), func(c *gin.Context) {
		user.ListUsers(c)
	})
	userGroup.GET("/info", utils.AuthMiddleware(), func(c *gin.Context) {
		user.GetUserInfo(c)
	})

	emailGroup := userGroup.Group("/email")

	emailGroup.POST("/bind",
		utils.AuthMiddleware(),
		utils.OPLoggerMiddleware("用户", "绑定邮箱"),
		func(c *gin.Context) {
			user.BindEmail(c)
		})
	emailGroup.POST("/verify",
		utils.AuthMiddleware(),
		utils.OPLoggerMiddleware("用户", "验证邮箱"),
		func(c *gin.Context) {
			user.VerifyEmail(c)
		})

	pswGroup := userGroup.Group("/psw")

	pswGroup.POST("/reset",
		func(c *gin.Context) {
			user.ResetPassword(c)
		})
	pswGroup.POST("/verify",
		func(c *gin.Context) {
			user.VerifyResetPasswordEmail(c)
		})
}
