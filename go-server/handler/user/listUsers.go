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
	err := db.Select("*").Find(&users).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of users",
			"detail": err.Error(),
			"code":   "501",
		})
		return
	}

	var usersRes []model.User
	for _, user := range users {
		usersRes = append(usersRes, user)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Get user list successfully",
		"rows":    usersRes,
		"code":    201,
	})
}
