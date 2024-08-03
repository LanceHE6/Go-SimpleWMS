package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRequest struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Login(context *gin.Context) {
	var data loginRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	account := data.Account
	password := data.Password

	db := my_db.GetMyDbConnection()
	var user model.User

	err := db.Where("(account=? and password=?) OR (email=? and password=?)", account, password, account, password).First(&user).Error

	if err != nil {
		context.JSON(http.StatusOK, response.Response(202, "Account or password is incorrect", nil))
		return
	} else {
		sessionID := utils.GenerateUUID(-1)
		token, err := utils.GenerateToken(user.Uid, user.Permission, user.CreatedAt.String(), sessionID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Cannot generate token", err.Error()))
			return
		}

		// token写入数据库
		err = db.Model(&user).Update("session_id", sessionID).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(502, "Cannot update session id", err.Error()))
			return
		}

		context.JSON(http.StatusOK, response.Response(201, "Login successfully", gin.H{
			"token": token,
			"user":  user,
		}))
	}

}
