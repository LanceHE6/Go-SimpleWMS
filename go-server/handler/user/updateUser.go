package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
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
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	uid := data.Uid
	password := data.Password
	nickname := data.Nickname
	permission := data.Permission
	phone := data.Phone

	//if password == "" && nickname == "" && permission == 0 && phone == "" {
	//	context.JSON(http.StatusBadRequest, gin.H{
	//		"message": "One of password, nickname, permission and phone is required",
	//		"code":    402,
	//	})
	//	return
	//}

	db := myDb.GetMyDbConnection()

	// 判断该用户是否已存在
	err := db.Model(&model.User{}).Where("uid=?", uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "User not found", nil))
		return
	}
	var updateData = map[string]interface{}{
		"nickname":   nickname,
		"permission": permission,
		"phone":      phone,
	}
	if password != "" {
		updateData["password"] = password
	}

	err = db.Model(&model.User{}).Where("uid = ?", uid).Updates(updateData).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Update user failed", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Update user successfully", nil))
}
