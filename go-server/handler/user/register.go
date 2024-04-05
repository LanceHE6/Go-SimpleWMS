package user

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	if tx == nil {
		return http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   "501",
		}
	}

	// 函数结束时回滚事务以关闭数据库连接，若已经提交则不会产生影响
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)

	// 锁定 user 表
	_, err = tx.Exec("SELECT * FROM user FOR UPDATE")
	if err != nil {
		return http.StatusInternalServerError, gin.H{
			"error":  "Cannot lock user table",
			"detail": err.Error(),
			"code":   507,
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

	newUid := "u" + utils.GenerateUuid(8)
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
