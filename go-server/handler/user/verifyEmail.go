package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type verifyEmailRequest struct {
	Uid   string `json:"uid" form:"uid" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	Code  string `json:"code" form:"code" binding:"required"`
}

func VerifyEmail(context *gin.Context) {
	var data verifyEmailRequest
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
	code := data.Code

	db := myDb.GetMyDbConnection()

	var verification model.VerificationCode

	// 查询数据库中是否存在该验证码
	notExist := db.Model(&model.VerificationCode{}).Where("uid = ? AND email = ? AND code = ? AND used = ?", uid, email, code, false).First(&verification).RecordNotFound()
	if notExist {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid verification code",
			"code":    401,
		})
		return
	}
	// 验证是否过期
	if time.Since(verification.CreatedAt) > 5*time.Minute {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The verification code has expired",
			"code":    402,
		})
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

	context.JSON(http.StatusOK, gin.H{
		"message": "Email verification successful",
		"code":    200,
	})
}
