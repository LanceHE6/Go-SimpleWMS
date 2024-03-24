package main

import (
	"Go_simpleWMS/handler/test"
	"Go_simpleWMS/handler/user"
	"Go_simpleWMS/handler/warehouse"
	"Go_simpleWMS/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(`
 ________  ___  _____ ______   ________  ___       _______   ___       __   _____ ______   ________      
|\   ____\|\  \|\   _ \  _   \|\   __  \|\  \     |\  ___ \ |\  \     |\  \|\   _ \  _   \|\   ____\     
\ \  \___|\ \  \ \  \\\__\ \  \ \  \|\  \ \  \    \ \   __/|\ \  \    \ \  \ \  \\\__\ \  \ \  \___|_    
 \ \_____  \ \  \ \  \\|__| \  \ \   ____\ \  \    \ \  \_|/_\ \  \  __\ \  \ \  \\|__| \  \ \_____  \   
  \|____|\  \ \  \ \  \    \ \  \ \  \___|\ \  \____\ \  \_|\ \ \  \|\__\_\  \ \  \    \ \  \|____|\  \  
    ____\_\  \ \__\ \__\    \ \__\ \__\    \ \_______\ \_______\ \____________\ \__\    \ \__\____\_\  \ 
   |\_________\|__|\|__|     \|__|\|__|     \|_______|\|_______|\|____________|\|__|     \|__|\_________\
   \|_________|                                                                              \|_________|
	`)
	utils.InitDB()
	defer utils.CloseDB()
	ginServer := gin.Default()
	// 解决跨域问题
	ginServer.Use(cors.Default())

	ginServer.GET("/ping", func(context *gin.Context) {
		test.Ping(context)
	})

	// 路由分组
	userGroup := ginServer.Group("/user")

	userGroup.POST("/register", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		user.Register(context)
	})
	userGroup.POST("/login", func(context *gin.Context) {
		user.Login(context)
	})
	userGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		user.DeleteUser(context)
	})
	userGroup.PUT("/update", utils.AuthMiddleware(), func(context *gin.Context) {
		user.UpdateUser(context)
	})

	warehouseGroup := ginServer.Group("/warehouse")

	warehouseGroup.POST("/add", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		warehouse.AddWarehouse(context)
	})

	warehouseGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		warehouse.DeleteWarehouse(context)
	})

	warehouseGroup.PUT("/update", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		warehouse.UpdateWarehouse(context)
	})

	warehouseGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		warehouse.ListWarehouse(context)
	})

	err := ginServer.Run(":8080")
	if err != nil {
		return
	}
}
