package main

import (
	"Go_simpleWMS/route"
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

	route.Route(ginServer)

	err := ginServer.Run(":8080")
	if err != nil {
		return
	}
}
