package test

import (
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"msg": "It's SimpleWMS!",
	})
}
