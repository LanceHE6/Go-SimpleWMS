package auth

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(context *gin.Context) {

	uid, _, createdAt, err := utils.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"code":    401,
		})
		return
	}
	// 判断是否在数据库中
	db := myDb.GetMyDbConnection()
	var user model.User
	err = db.Where("uid=? and created_at=?", uid, createdAt).First(&user).Error

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"code":    101,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 201,
		"uid":  uid,
	})
}
