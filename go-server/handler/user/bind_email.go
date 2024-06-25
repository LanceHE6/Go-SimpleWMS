package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type bindRequest struct {
	Uid   string `json:"uid" form:"uid" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
}

// BindEmail 发送绑定邮箱验证码
func BindEmail(context *gin.Context) {
	var data bindRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}

	uid := data.Uid
	email := data.Email

	db := my_db.GetMyDbConnection()
	var user model.User

	// 查询用户是否存在
	notExist := db.Model(model.User{}).Where("uid = ?", uid).First(&user).RecordNotFound()
	if notExist {
		context.JSON(http.StatusOK, response.Response(402, "User not found", nil))
		return
	}

	code := utils.CreateVerifyCode()

	// 使用通道传递错误信息
	errorChan := make(chan error)
	go func() {
		err := utils.SendEmail(email, user.Account, code, utils.BindEmail)
		errorChan <- err
	}()

	var verification = model.VerificationCode{
		Id:        uid,
		Email:     email,
		Code:      code,
		CreatedAt: time.Now(),
		Used:      false,
	}
	// 判断是否有关于该用户的记录
	var verificationCode model.VerificationCode
	notExist = db.Model(model.VerificationCode{}).Where("id = ?", uid).First(&verificationCode).RecordNotFound()
	if !notExist {
		var updateData = map[string]interface{}{
			"code":       code,
			"created_at": time.Now(),
			"used":       false,
			"email":      email,
		}
		db.Model(model.VerificationCode{}).Where("id = ?", uid).Updates(updateData)
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

type verifyEmailRequest struct {
	Uid   string `json:"uid" form:"uid" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	Code  string `json:"code" form:"code" binding:"required"`
}

// VerifyEmail 校验验证码，绑定邮箱
func VerifyEmail(context *gin.Context) {
	var data verifyEmailRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	uid := data.Uid
	email := data.Email
	code := data.Code

	db := my_db.GetMyDbConnection()

	var verification model.VerificationCode

	// 查询数据库中是否存在该验证码
	notExist := db.Model(&model.VerificationCode{}).Where("id = ? AND email = ? AND code = ? AND used = ?", uid, email, code, false).First(&verification).RecordNotFound()
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

	// 更新用户邮箱
	var user model.User
	db.Model(&model.User{}).Where("uid = ?", uid).First(&user)
	user.Email = email
	db.Save(&user)

	context.JSON(http.StatusOK, response.Response(201, "Email verification successful", nil))
}
