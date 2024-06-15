package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUsers(context *gin.Context) {
	db := myDb.GetMyDbConnection()

	var users []model.User
	err := db.Select("*").Find(&users).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Get user list failed", err.Error()))
		return
	}

	var usersRes []model.User
	for _, user := range users {
		usersRes = append(usersRes, user)
	}

	context.JSON(http.StatusOK, response.Response(200, "Get user list success", gin.H{
		"rows": usersRes,
	}))
}
