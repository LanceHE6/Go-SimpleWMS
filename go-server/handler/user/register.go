package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type registerRequest struct {
	Account    string `json:"account" form:"account" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	Permission int    `json:"permission" form:"permission" binding:"required"`
	NickName   string `json:"nickname" form:"nickname" binding:"required"`
	Phone      string `json:"phone" form:"phone"`
}

func Register(context *gin.Context) {
	var data registerRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}

	status, returnInfo := DoRegister(data)
	context.JSON(status, returnInfo)

}

// DoRegister 执行注册逻辑函数
func DoRegister(userData registerRequest) (int, gin.H) {
	account := userData.Account
	password := userData.Password
	permission := userData.Permission
	nickName := userData.NickName
	phone := userData.Phone

	db := my_db.GetMyDbConnection()

	// 判断该账户是否已被注册
	var user model.User
	if err := db.Where("account = ?", account).First(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, response.Response(202, fmt.Sprintf("The account '%s' has been registered", account), nil)
	}

	// 插入新用户
	newUid := "u" + utils.GenerateUuid(8)
	user = model.User{
		Uid:        newUid,
		Account:    account,
		Password:   password,
		Nickname:   nickName,
		Permission: permission,
		Phone:      phone,
	}
	if err := db.Create(&user).Error; err != nil {
		return http.StatusInternalServerError, response.ErrorResponse(501, "Failed to register user", err.Error())
	}

	return http.StatusOK, response.Response(201, fmt.Sprintf("Successfully registered user '%s'", account), gin.H{
		"uid": newUid,
	})
}
