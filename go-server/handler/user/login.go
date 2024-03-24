package user

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	account := context.PostForm("account")
	password := context.PostForm("password")

	if account == "" || password == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Account and password are required"})
		return

	}
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
