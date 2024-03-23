package handler

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Register(context *gin.Context) {
	account := context.PostForm("account")
	password := context.PostForm("password")
	permission := context.PostForm("permission")
	nickName := context.PostForm("nick_name")

	if account == "" || password == "" || permission == "" || nickName == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Account, password, permission and nick_name are required"})
		return
	}

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 判断该账户是否已被注册
	var registered int
	err = tx.QueryRow("SELECT count(account) FROM user WHERE account=?", account).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the number of users for this account"})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{"message": "The account has been registered"})
		return
	}
	// 获取最近注册的用户的 uid
	var lastUid string
	err = tx.QueryRow("SELECT uid FROM user ORDER BY register_time DESC LIMIT 1").Scan(&lastUid)
	// 如果没有用户，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastUid = "00000000"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get last uid"})
		return
	}

	// 增加最近注册的用户的 uid
	nextUid, err := strconv.Atoi(lastUid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot convert uid to integer"})
		return
	}
	nextUid++
	newUid := fmt.Sprintf("%08d", nextUid) // 转换为 8 位字符串

	// 获取当前时间戳
	registerTime := time.Now().Unix()

	// 插入新用户
	_, err = tx.Exec("INSERT INTO user (uid, account, password, nick_name, permission, register_time) VALUES (?, ?, ?, ?, ?, ?)",
		newUid, account, password, nickName, permission, registerTime)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot insert new user"})
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit transaction"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"uid":     newUid,
	})
}

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

func DeleteUser(context *gin.Context) {
	uid := context.PostForm("uid")
	if uid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "UID is required"})
		return
	}
	tx, err := utils.GetDbConnection()
	// 开始一个新的事务
	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}
	// 删除用户
	_, err = tx.Exec("DELETE FROM user WHERE uid=?", uid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete user"})
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func UpdateUser(context *gin.Context) {
	uid := context.PostForm("uid")
	password := context.PostForm("password")
	nickName := context.PostForm("nick_name")
	permission := context.PostForm("permission")

	if uid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "UID is required"})
		return
	}

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	if password == "" && nickName == "" && permission == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "At least one of password, nick_name and permission is required"})
		return
	}
	// 拼接sql语句
	updateSql := "UPDATE user SET "
	if password != "" {
		updateSql += "password='" + password + "',"
	}
	if nickName != "" {
		updateSql += "nick_name='" + nickName + "',"
	}
	if permission != "" {
		updateSql += "permission=" + permission + ","
	}
	updateSql = updateSql[:len(updateSql)-1] // 去掉最后一个逗号
	updateSql += " WHERE uid='" + uid + "'"
	_, err = tx.Exec(updateSql)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update user"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}
