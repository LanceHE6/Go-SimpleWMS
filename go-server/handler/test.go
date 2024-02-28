package handler

import (
	"github.com/gin-gonic/gin"
)

func Test(context *gin.Context) {
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
