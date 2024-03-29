package auth

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthByHeader(context *gin.Context) {

	uid, err := utils.GetUidByContext(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"code":    401,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 201,
		"uid":  uid,
	})
}
