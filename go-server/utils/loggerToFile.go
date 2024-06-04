package utils

import (
	"Go_simpleWMS/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.ServerConfig.SERVER.LOG.PATH
	maxLogFiles := config.ServerConfig.SERVER.LOG.MAX_FILES

	var mutex sync.Mutex // 使用互斥锁来确保并发安全
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
		// 在写入日志前，检查并删除旧日志
		mutex.Lock()
		defer mutex.Unlock()
		checkAndDeleteOldLogs(logFilePath, maxLogFiles)
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

// checkAndDeleteOldLogs 检查并删除旧的日志文件
func checkAndDeleteOldLogs(logPath string, maxFiles int) {
	files, err := ioutil.ReadDir(logPath)
	if err != nil {
		fmt.Println("Failed to read log directory:", err)
		return
	}

	// 过滤出以日期命名的日志文件
	var logFiles []os.FileInfo
	for _, file := range files {
		if strings.HasSuffix(file.Name(), "-server.log") {
			logFiles = append(logFiles, file)
		}
	}

	// 根据文件名（日期）对日志文件进行排序
	sort.Slice(logFiles, func(i, j int) bool {
		return parseLogFileDate(logFiles[i].Name()).Before(parseLogFileDate(logFiles[j].Name()))
	})

	// 删除超过最大数量的旧日志文件
	for len(logFiles) > maxFiles {
		oldestFile := logFiles[0]
		oldestFilePath := filepath.Join(logPath, oldestFile.Name())
		err := os.Remove(oldestFilePath)
		if err != nil {
			fmt.Printf("Failed to delete oldest log file %s: %v\n", oldestFilePath, err)
			continue
		}
		logFiles = logFiles[1:] // 移除已删除的文件
	}
}

// parseLogFileDate 解析日志文件名中的日期部分
func parseLogFileDate(fileName string) time.Time {
	// 文件名格式为 "2023-01-02-server.log"
	parts := strings.Split(fileName, "-")
	if len(parts) < 2 {
		return time.Time{} // 返回一个零值时间
	}
	dateStr := parts[0]
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{} // 返回一个零值时间
	}
	return t
}
