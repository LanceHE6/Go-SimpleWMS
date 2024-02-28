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
	ginServer.POST("/register", func(context *gin.Context) {
		handler.Register(context)
	})

	err := ginServer.Run(":8080")
	if err != nil {
		return
	}
}
