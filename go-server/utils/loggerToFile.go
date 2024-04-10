package utils

import (
	"Go_simpleWMS/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.ServerConfig.SERVER.LOG.PATH
	return func(c *gin.Context) {
		// 获取当前日期
		date := time.Now().Format("2006-01-02")
		// 创建日志文件名
		logFileName := fmt.Sprintf("%s-server.log", date)
		//日志文件路径
		fileName := path.Join(logFilePath, logFileName)

		// 创建日志文件所在的目录
		if err := os.MkdirAll(logFilePath, 0755); err != nil {
			fmt.Println("Failed to create log file directory:", err)
			return
		}
		//写入文件
		src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Failed to open log file:", err)
			return
		}
		defer func(src *os.File) {
			err := src.Close()
			if err != nil {
				_ = fmt.Errorf("can not close log file")
				return
			}
		}(src)

		//实例化
		logger := logrus.New()
		//设置输出
		logger.Out = src
		//设置日志级别
		logger.SetLevel(logrus.DebugLevel)
		//设置日志格式
		logger.SetFormatter(&logrus.TextFormatter{})

		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
