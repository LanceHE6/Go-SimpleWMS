package staff

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

type addStaffRequest struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Phone  string `json:"phone" form:"phone"`
	DeptId string `json:"dept_id" form:"dept_id" binding:"required"`
}

func AddStaff(context *gin.Context) {
	var data addStaffRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Staff name or DeptId is required"})
		return
	}
	staffName := data.Name
	staffDeptId := data.DeptId
	phone := data.Phone

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
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
		})
		return
	}

	// 判断该仓库是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM staff WHERE name=?", staffName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of staffs for this staff name",
			"detail": err.Error(),
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{"message": "The staff already exists"})
		return
	}

	// 获取最近注册的员工的 sid
	var lastSid string
	err = tx.QueryRow("SELECT sid FROM staff ORDER BY add_time DESC LIMIT 1").Scan(&lastSid)
	// 如果没有用户，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastSid = "s00000000"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get last Sid",
			"detail": err.Error(),
		})
		return
	}
	lastSid = lastSid[1:]
	// 增加最近注册的用户的 uid
	nextSid, err := strconv.Atoi(lastSid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot convert Sid to integer",
			"detail": err.Error(),
		})
		return
	}
	nextSid++
	newSid := fmt.Sprintf("s%08d", nextSid) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	// 增加仓库
	_, err = tx.Exec("INSERT INTO staff(sid, name, add_time, phone, department) VALUES(?, ?, ?, ?, ?)", newSid, staffName, addTime, phone, staffDeptId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the staff",
			"detail": err.Error(),
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Staff added successfully"})
}
