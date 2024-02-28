package handler

import (
	"Go_simpleWMS/utils"
	"database/sql"
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

	db := utils.GetDbConnection()

	// 开始一个新的事务
	tx, err := db.Begin()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {

		}
	}(tx) // 如果出错，回滚事务

	// 获取最近注册的用户的 uid
	var lastUid string
	err = tx.QueryRow("SELECT uid FROM user ORDER BY register_time DESC LIMIT 1").Scan(&lastUid)
	if err != nil {
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
	_, err = tx.Exec("INSERT INTO user (uid, account, password, permission, register_time) VALUES (?, ?, ?, ?, ?)",
		newUid, account, password, permission, registerTime)
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
