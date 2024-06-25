package utils

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// OPLoggerMiddleware 操作日志中间件
func OPLoggerMiddleware(resource string, operation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID，这里假设从请求头中获取
		uid, _, _, _ := GetUserInfoByContext(c)

		// 处理请求
		c.Next()

		// 获取状态和错误信息
		status := "成功"
		if c.Writer.Status() != http.StatusOK {
			status = "失败"
		}

		// 记录日志
		opLog := model.OPLog{
			Uid:       uid,
			Resource:  resource,
			Operation: operation,
			Status:    status,
		}
		db := my_db.GetMyDbConnection()
		err := db.Create(&opLog).Error

		if err != nil {
			log.Println("Failed to insert log:", err)
		}
	}
}
