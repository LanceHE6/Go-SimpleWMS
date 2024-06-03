package main

import (
	"Go_simpleWMS/config"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/route"
	"Go_simpleWMS/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/semaphore"
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

	ginServer := gin.Default()
	// 解决跨域问题
	ginServer.Use(utils.Cors())

	// 连接数据库
	myDb.Init()
	defer myDb.CloseMyDb()

	// 设置日志等级
	if config.ServerConfig.SERVER.MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 日志到文件
	ginServer.Use(utils.LoggerToFile())

	sem := semaphore.NewWeighted(50) // 最大并发处理数为50

	route.Route(ginServer, sem)

	err := ginServer.Run(":" + config.ServerConfig.SERVER.PORT)
	if err != nil {
		return
	}
}
