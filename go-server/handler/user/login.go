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
		context.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": "Account and password are required"}})
		return
	}
	account := data.Account
	password := data.Password

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	var uid string
	var permission int
	err = tx.QueryRow("SELECT uid, permission FROM user WHERE account = ? AND password = ?", account, password).Scan(&uid, &permission)
	if err != nil {
		context.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Incorrect account or password"})
		return
	} else {
		token, err := utils.GenerateToken(uid, permission)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot generate token"})
			return
		}

		// token写入数据库
		_, err = tx.Exec("UPDATE user set token=? WHERE uid=?", token, uid)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update token"})
			return
		}

		// 提交事务
		err = tx.Commit()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit transaction"})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
			"token":   token})
	}

}
