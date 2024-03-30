package user

import (
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
			"message": "Account and password are required",
			"code":    401,
		})
		return
	}
	account := data.Account
	password := data.Password

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	var uid string
	var permission int
	var registerTime string
	err = tx.QueryRow("SELECT uid, permission, register_time FROM user WHERE account = ? AND password = ?", account, password).Scan(&uid, &permission, &registerTime)
	if err != nil {
		context.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Incorrect account or password"})
		return
	} else {
		token, err := utils.GenerateToken(uid, permission, registerTime)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot generate token",
				"detail": err.Error(),
				"code":   502,
			})
			return
		}

		// token写入数据库
		_, err = tx.Exec("UPDATE user set token=? WHERE uid=?", token, uid)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot update token",
				"detail": err.Error(),
				"code":   503,
			})
			return
		}

		// 提交事务
		err = tx.Commit()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot commit transaction",
				"detail": err.Error(),
				"code":   504,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
			"token":   token,
			"code":    201,
		})
	}

}
