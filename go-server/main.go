package main

import (
	"Go_simpleWMS/handler"
	_ "Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

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

	err := ginServer.Run(":8080")
	if err != nil {
		return
	}
}
