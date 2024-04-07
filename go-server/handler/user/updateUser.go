package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type updateRequest struct {
	Uid        string `json:"uid" form:"uid" binding:"required"`
	Password   string `json:"password" form:"password"`
	Nickname   string `json:"nickname" form:"nickname"`
	Permission int    `json:"permission" form:"permission"`
	Phone      string `json:"phone" form:"phone"`
}

func UpdateUser(context *gin.Context) {
	var data updateRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	uid := data.Uid
	password := data.Password
	nickname := data.Nickname
	permission := data.Permission
	phone := data.Phone

	if password == "" && nickname == "" && permission == 0 && phone == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "One of password, nickname, permission and phone is required",
			"code":    402,
		})
		return
	}

	db := myDb.GetMyDbConnection()

	// 判断该用户是否已存在
	err := db.Model(&model.User{}).Where("uid=?", uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The user does not exist",
			"code":    403,
		})
		return
	}

	user := model.User{
		Uid:        uid,
		Password:   password,
		Nickname:   nickname,
		Permission: permission,
		Phone:      phone,
	}

	err = db.Model(&user).Where("uid = ?", user.Uid).Updates(user).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update user",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"code":    201,
	})
}
