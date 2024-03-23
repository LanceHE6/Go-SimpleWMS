package main

import (
	"Go_simpleWMS/handler"
	"Go_simpleWMS/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitDB()
	defer utils.CloseDB()
	ginServer := gin.Default()
	// 解决跨域问题
	ginServer.Use(cors.Default())

	ginServer.GET("/ping", func(context *gin.Context) {
		handler.Test(context)
	})

	// 路由分组
	userGroup := ginServer.Group("/user")

	userGroup.POST("/register", func(context *gin.Context) {
		handler.Register(context)
	})
	userGroup.POST("/login", func(context *gin.Context) {
		handler.Login(context)
	})
	userGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		handler.DeleteUser(context)
	})
	userGroup.PUT("/update", utils.AuthMiddleware(), func(context *gin.Context) {
		handler.UpdateUser(context)
	})

	warehouseGroup := ginServer.Group("/warehouse")

	warehouseGroup.POST("/add", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		handler.AddWarehouse(context)
	})

	warehouseGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		handler.DeleteWarehouse(context)
	})

	err := ginServer.Run(":8080")
	if err != nil {
		return
	}
}
