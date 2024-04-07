package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUsers(context *gin.Context) {
	db := myDb.GetMyDbConnection()

	var users []model.User
	err := db.Select([]string{"uid", "account", "permission", "created_at", "phone", "nickname"}).Find(&users).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of users",
			"detail": err.Error(),
			"code":   "501",
		})
		return
	}

	var usersRes []gin.H
	for _, user := range users {
		userRes := gin.H{
			"uid":        user.Uid,
			"account":    user.Account,
			"permission": user.Permission,
			"created_at": user.CreatedAt,
			"phone":      user.Phone,
			"nickname":   user.Nickname,
		}
		usersRes = append(usersRes, userRes)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Get user list successfully",
		"rows":    usersRes,
		"code":    201,
	})
}
