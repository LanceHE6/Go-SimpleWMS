package user

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
	phone := context.PostForm("phone")

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
	_, err = tx.Exec("INSERT INTO user (uid, account, password, nick_name, permission, register_time, phone) VALUES (?, ?, ?, ?, ?, ?, ?)",
		newUid, account, password, nickName, permission, registerTime, phone)
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
