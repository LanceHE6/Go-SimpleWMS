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

type registerRequest struct {
	Account    string `json:"account" form:"account" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	Permission int    `json:"permission" form:"permission" binding:"required"`
	NickName   string `json:"nick_name" form:"nick_name" binding:"required"`
	Phone      string `json:"phone" form:"phone"`
}

func Register(context *gin.Context) {
	var data registerRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Account, password, permission and nick_name are required",
			"code":    401,
		})
		return
	}
	account := data.Account
	password := data.Password
	permission := data.Permission
	nickName := data.NickName
	phone := data.Phone

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   "501",
		})
		return
	}

	// 判断该账户是否已被注册
	var registered int
	err = tx.QueryRow("SELECT count(account) FROM user WHERE account=?", account).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of users for this account",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The account has been registered",
			"code":    402,
		})
		return
	}
	// 获取最近注册的用户的 uid
	var lastUid string
	err = tx.QueryRow("SELECT uid FROM user ORDER BY register_time DESC LIMIT 1").Scan(&lastUid)
	// 如果没有用户，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastUid = "u00000000"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get last uid",
			"detail": err.Error(),
			"code":   503,
		})
		return
	}
	lastUid = lastUid[1:]
	// 增加最近注册的用户的 uid
	nextUid, err := strconv.Atoi(lastUid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot convert uid to integer",
			"detail": err.Error(),
			"code":   504,
		})
		return
	}
	nextUid++
	newUid := fmt.Sprintf("u%08d", nextUid) // 转换为 8 位字符串

	// 获取当前时间戳
	registerTime := time.Now().Unix()

	// 插入新用户
	_, err = tx.Exec("INSERT INTO user (uid, account, password, nick_name, permission, register_time, phone) VALUES (?, ?, ?, ?, ?, ?, ?)",
		newUid, account, password, nickName, permission, registerTime, phone)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert new user",
			"detail": err.Error(),
			"code":   505,
		})
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit transaction",
			"detail": err.Error(),
			"code":   506,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"uid":     newUid,
		"code":    201,
	})
}
