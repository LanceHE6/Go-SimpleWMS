package auth

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthByHeader(context *gin.Context) {

	uid, _, registerTime, err := utils.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"code":    401,
		})
		return
	}
	// 判断是否在数据库中
	tx, _ := utils.GetDbConnection()
	var isExist int
	err = tx.QueryRow("SELECT count(*) from user where uid=? and register_time=?", uid, registerTime).Scan(&isExist)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of uid for this uid",
			"detail": err.Error(),
			"code":   501,
		})
		context.Abort()
		return
	}
	if isExist <= 0 {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"code":    101,
		})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 201,
		"uid":  uid,
	})
}
