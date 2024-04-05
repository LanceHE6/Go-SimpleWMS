package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/semaphore"
	"net/http"
)

func SemaphoreMiddleware(sem *semaphore.Weighted) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试获取一个信号量
		if err := sem.Acquire(c.Request.Context(), 1); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to acquire semaphore"})
			c.Abort()
			return
		}

		// 确保在请求完成后释放信号量
		defer sem.Release(1)

		// 继续处理请求
		c.Next()
	}
}
