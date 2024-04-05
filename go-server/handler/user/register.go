package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
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

	db := myDb.GetMyDbConnection()

	// 判断该账户是否已被注册
	var user model.User
	if err := db.Where("account = ?", account).First(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusForbidden, gin.H{
			"message": fmt.Sprintf("The account '%s' has been registered", account),
			"code":    402,
		}
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
		return http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert new user",
			"detail": err.Error(),
			"code":   505,
		}
	}

	return http.StatusOK, gin.H{
		"message": "User registered successfully",
		"uid":     newUid,
		"code":    201,
	}
}
