package auth

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthByHeader(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")

	uid, err := utils.GetUidByToken(bearerToken[1])
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"uid": uid})
}
