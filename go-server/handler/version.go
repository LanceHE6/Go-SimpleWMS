package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const version = "v0.0.3.20240607_Alpha"

func GetVersion(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Hello Go-SimpleWMS",
		"data": gin.H{
			"version": version,
		},
	})
}
