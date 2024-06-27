package handler

import (
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

const version = "v0.6.13.20240627_Alpha"

func GetVersion(context *gin.Context) {
	context.JSON(http.StatusOK, response.Response(200, "Hello Simple-WMS", gin.H{
		"version": version,
	}))
}
