package department

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

type addDepartmentRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func AddDepartment(context *gin.Context) {
	var data addDepartmentRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	depName := data.Name

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	// 判断该部门是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM department WHERE name=?", depName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of department for this department name",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The department name already exists",
			"code":    402,
		})
		return
	}

	// 获取最近注册的部门的 did
	var lastDid string
	err = tx.QueryRow("SELECT did FROM department ORDER BY add_time DESC LIMIT 1").Scan(&lastDid)
	// 如果没有用户，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastDid = "d0000"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get last Did",
			"detail": err.Error(),
			"code":   503,
		})
		return
	}
	lastDid = lastDid[1:]
	nextDid, err := strconv.Atoi(lastDid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot convert Did to integer",
			"detail": err.Error(),
			"code":   504,
		})
		return
	}
	nextDid++
	newDid := fmt.Sprintf("d%04d", nextDid) // 转换为 4 位字符串

	addTime := time.Now().Unix()
	// 增加仓库
	_, err = tx.Exec("INSERT INTO department(did, name, add_time) VALUES(?, ?, ?)", newDid, depName, addTime)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the department",
			"detail": err.Error(),
			"code":   505,
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
			"code":   506,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Department added successfully",
		"code":    201,
	})
}
