package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteRequest struct {
	Uid string `json:"uid" form:"uid" binding:"required"`
}

func DeleteUser(context *gin.Context) {
	var data deleteRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	uid := data.Uid

	targetUid, _, _, err := utils.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"code":    101,
		})
		return
	}
	if targetUid == uid {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid target uid",
			"code":    402,
		})
		return
	}
	db := myDb.GetMyDbConnection()

	// 删除用户
	err = db.Delete(&model.User{}, "uid=?", uid).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete user",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"code":    201,
	})
}
