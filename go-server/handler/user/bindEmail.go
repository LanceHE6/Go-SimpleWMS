package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type bindRequest struct {
	Uid   string `json:"uid" form:"uid" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
}

func BindEmail(context *gin.Context) {
	var data bindRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}

	uid := data.Uid
	email := data.Email

	db := myDb.GetMyDbConnection()
	var user model.User

	// 查询用户是否存在
	notExist := db.Model(model.User{}).Where("uid = ?", uid).First(&user).RecordNotFound()
	if notExist {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "User does not exist",
			"code":    402,
		})
		return
	}

	code := utils.CreateVerifyCode()

	// 使用通道传递错误信息
	errorChan := make(chan error)
	go func() {
		err := utils.SendVerifyEmail(email, user.Account, code)
		errorChan <- err
	}()

	var verification = model.VerificationCode{
		Uid:       uid,
		Email:     email,
		Code:      code,
		CreatedAt: time.Now(),
		Used:      false,
	}
	// 判断是否有关于该用户的记录
	var verificationCode model.VerificationCode
	notExist = db.Model(model.VerificationCode{}).Where("uid = ?", uid).First(&verificationCode).RecordNotFound()
	if !notExist {
		var updateData = map[string]interface{}{
			"code":       code,
			"created_at": time.Now(),
			"used":       false,
			"email":      email,
		}
		db.Model(model.VerificationCode{}).Where("uid = ?", uid).Updates(updateData)
	} else {
		db.Create(&verification)
	}

	// 处理协程中的错误信息
	if err := <-errorChan; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to send email",
			"code":    501,
			"detail":  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Verification code sent successfully",
		"code":    200,
	})
}
