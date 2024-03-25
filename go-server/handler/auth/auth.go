package auth

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthByHeader(context *gin.Context) {

	uid, err := utils.GetUidByContext(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"uid": uid})
}
