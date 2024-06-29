package main

import (
	"Go_simpleWMS/config"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/router"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/semaphore"
)

func main() {
	utils.LogoPrint()

	ginServer := gin.Default()
	// 解决跨域问题
	ginServer.Use(utils.Cors())

	// 连接数据库
	my_db.Init()
	defer my_db.CloseMyDb()

	// 设置日志等级
	if config.ServerConfig.SERVER.MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 日志到文件
	ginServer.Use(utils.LoggerToFile())

	sem := semaphore.NewWeighted(50) // 最大并发处理数为50

	router.Route(ginServer, sem)

	err := ginServer.Run(":" + config.ServerConfig.SERVER.PORT)
	if err != nil {
		return
	}
}
