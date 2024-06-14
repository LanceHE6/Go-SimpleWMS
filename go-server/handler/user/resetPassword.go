package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type resetPswRequest struct {
	Account string `json:"account" form:"account" binding:"required"`
	Email   string `json:"email" form:"email" binding:"required"`
}

// ResetPassword 发送重置密码验证码
func ResetPassword(context *gin.Context) {
	var data resetPswRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}

	account := data.Account
	email := data.Email

	db := myDb.GetMyDbConnection()
	var user model.User

	// 查询用户是否存在
	notExist := db.Model(model.User{}).Where("account = ? AND email = ?", account, email).First(&user).RecordNotFound()
	if notExist {
		context.JSON(http.StatusOK, response.Response(402, "User not found", nil))
		return
	}

	code := utils.CreateVerifyCode()

	// 使用通道传递错误信息
	errorChan := make(chan error)
	go func() {
		err := utils.SendEmail(email, user.Account, code, utils.ResetPasswordEmail)
		errorChan <- err
	}()

	var verification = model.VerificationCode{
		Id:        account,
		Email:     email,
		Code:      code,
		CreatedAt: time.Now(),
		Used:      false,
	}
	// 判断是否有关于该用户的记录
	var verificationCode model.VerificationCode
	notExist = db.Model(model.VerificationCode{}).Where("id = ?", account).First(&verificationCode).RecordNotFound()
	if !notExist {
		var updateData = map[string]interface{}{
			"code":       code,
			"created_at": time.Now(),
			"used":       false,
			"email":      email,
		}
		db.Model(model.VerificationCode{}).Where("id = ?", account).Updates(updateData)
	} else {
		db.Create(&verification)
	}

	// 处理协程中的错误信息
	if err := <-errorChan; err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to send verification code", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Success", nil))
}

type verifyResetPswRequest struct {
	Account     string `json:"account" form:"account" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	Code        string `json:"code" form:"code" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

// VerifyResetPasswordEmail 校验验证码，重置密码
func VerifyResetPasswordEmail(context *gin.Context) {
	var data verifyResetPswRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	account := data.Account
	email := data.Email
	code := data.Code
	newPassword := data.NewPassword

	db := myDb.GetMyDbConnection()

	var verification model.VerificationCode

	// 查询数据库中是否存在该验证码
	notExist := db.Model(&model.VerificationCode{}).Where("id = ? AND email = ? AND code = ? AND used = ?", account, email, code, false).First(&verification).RecordNotFound()
	if notExist {
		context.JSON(http.StatusOK, response.Response(202, "The verification code is invalid", nil))
		return
	}
	// 验证是否过期
	if time.Since(verification.CreatedAt) > 5*time.Minute {
		context.JSON(http.StatusOK, response.Response(203, "The verification code has expired", nil))
		return
	}
	// 更新验证码状态
	verification.Used = true
	db.Save(&verification)

	// 更新用户密码
	var user model.User
	db.Model(&model.User{}).Where("account = ? AND email = ?", account, email).First(&user)
	user.Password = newPassword
	db.Save(&user)

	context.JSON(http.StatusOK, response.Response(201, "Password reset successful", nil))
}
