package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const version = "v0.0.4.20240610_Alpha"

func GetVersion(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Hello Go-SimpleWMS",
		"data": gin.H{
			"version": version,
		},
	})
}
