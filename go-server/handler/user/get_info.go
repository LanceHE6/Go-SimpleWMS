package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(context *gin.Context) {
	uid, _, _, _ := utils.GetUserInfoByContext(context)
	db := my_db.GetMyDbConnection()
	var user model.User
	if err := db.Model(model.User{}).Where("uid = ?", uid).First(&user).RecordNotFound(); err {
		// 用户不存在
		context.JSON(http.StatusOK, response.Response(202, "The user does not exist", nil))
		return
	}

	context.JSON(http.StatusOK, response.Response(201, "success", gin.H{
		"user": user,
	}))

}
