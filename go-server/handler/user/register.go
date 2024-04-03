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
	NickName   string `json:"nickname" form:"nickname" binding:"required"`
	Phone      string `json:"phone" form:"phone"`
}

func Register(context *gin.Context) {
	var data registerRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}

	status, returnInfo := DoRegister(data)
	context.JSON(status, returnInfo)

}

// DoRegister 执行注册逻辑函数
func DoRegister(userData registerRequest) (int, gin.H) {
	account := userData.Account
	password := userData.Password
	permission := userData.Permission
	nickName := userData.NickName
	phone := userData.Phone

	tx, err := utils.GetDbConnection()
	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()

	if tx == nil {
		return http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   "501",
		}
	}

	// 判断该账户是否已被注册
	var registered int
	err = tx.QueryRow("SELECT count(account) FROM user WHERE account=?", account).Scan(&registered)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of users for this account",
			"detail": err.Error(),
			"code":   502,
		}
	}
	if registered >= 1 {
		return http.StatusForbidden, gin.H{
			"message": fmt.Sprintf("The account '%s' has been registered", account),
			"code":    402,
		}
	}
	// 获取最近注册的用户的 uid
	var lastUid string
	err = tx.QueryRow("SELECT uid FROM user ORDER BY register_time DESC LIMIT 1").Scan(&lastUid)
	// 如果没有用户，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastUid = "u00000000"
	} else if err != nil {
		return http.StatusInternalServerError, gin.H{
			"error":  "Cannot get last uid",
			"detail": err.Error(),
			"code":   503,
		}
	}
	lastUid = lastUid[1:]
	// 增加最近注册的用户的 uid
	nextUid, err := strconv.Atoi(lastUid)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"error":  "Cannot convert uid to integer",
			"detail": err.Error(),
			"code":   504,
		}
	}
	nextUid++
	newUid := fmt.Sprintf("u%08d", nextUid) // 转换为 8 位字符串

	// 获取当前时间戳
	registerTime := time.Now().Unix()

	// 插入新用户
	_, err = tx.Exec("INSERT INTO user (uid, account, password, nickname, permission, register_time, phone) VALUES (?, ?, ?, ?, ?, ?, ?)",
		newUid, account, password, nickName, permission, registerTime, phone)
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert new user",
			"detail": err.Error(),
			"code":   505,
		}
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit transaction",
			"detail": err.Error(),
			"code":   506,
		}
	}

	return http.StatusOK, gin.H{
		"message": "User registered successfully",
		"uid":     newUid,
		"code":    201,
	}
}
