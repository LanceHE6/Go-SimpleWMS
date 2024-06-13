package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	account := data.Account
	password := data.Password

	db := myDb.GetMyDbConnection()
	var user model.User

	err := db.Where("account=? and password=?", account, password).First(&user).Error

	if err != nil {
		context.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Incorrect account or password"})
		return
	} else {
		token, err := utils.GenerateToken(user.Uid, user.Permission, user.CreatedAt.String())
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot generate token",
				"detail": err.Error(),
				"code":   501,
			})
			return
		}

		// token写入数据库
		err = db.Model(&user).Update("token", token).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot update token",
				"detail": err.Error(),
				"code":   502,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
			"data": gin.H{
				"token": token,
				"user":  user,
			},
			"code": 201,
		})
	}

}
