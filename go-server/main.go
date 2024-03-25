package main

import (
	"Go_simpleWMS/handler/auth"
	"Go_simpleWMS/handler/goodsType"
	"Go_simpleWMS/handler/test"
	"Go_simpleWMS/handler/upload"
	"Go_simpleWMS/handler/user"
	"Go_simpleWMS/handler/warehouse"
	"Go_simpleWMS/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
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
	ginServer.Use(cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type", "Access-Token"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	}))

	ginServer.GET("/ping", func(context *gin.Context) {
		test.Ping(context)
	})
	// 鉴权接口
	ginServer.GET("/auth", utils.AuthMiddleware(), func(context *gin.Context) {
		auth.AuthByHeader(context)
	})
	ginServer.POST("/upload", func(context *gin.Context) {
		upload.UploadFile(context)
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
	userGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		user.ListUsers(context)
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

	goodsTypeGroup := ginServer.Group("/gt")

	goodsTypeGroup.POST("/add", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		goodsType.AddGoodsType(context)
	})

	goodsTypeGroup.PUT("/update", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		goodsType.UpdateGoodsType(context)
	})

	goodsTypeGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		goodsType.DeleteGoodsType(context)
	})

	goodsTypeGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		goodsType.ListGoodsType(context)
	})

	err := ginServer.Run(":8080")
	if err != nil {
		return
	}
}
