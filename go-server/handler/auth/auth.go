package auth

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(context *gin.Context) {

	uid, _, createdAt, err := utils.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, response.Response(401, "Invalid token", nil))
		return
	}
	// 判断是否在数据库中
	db := my_db.GetMyDbConnection()
	var user model.User
	err = db.Where("uid=? and created_at=?", uid, createdAt).First(&user).Error

	if err != nil {
		context.JSON(http.StatusUnauthorized, response.Response(402, "Invalid token", nil))
		return
	}
	context.JSON(http.StatusOK, response.Response(201, "Hello", gin.H{
		"uid": uid,
	}))
}
