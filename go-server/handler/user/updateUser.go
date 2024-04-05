package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
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

	db := myDb.GetMyDbConnection()

	user := model.User{
		Uid:        data.Uid,
		Password:   data.Password,
		Nickname:   data.Nickname,
		Permission: data.Permission,
		Phone:      data.Phone,
	}

	err := db.Model(&user).Where("uid = ?", user.Uid).Updates(user).Error
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
