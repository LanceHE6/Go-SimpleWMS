package unit

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

type addUnitRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func AddUnit(context *gin.Context) {
	var data addUnitRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unit name is required " + err.Error(),
			"code":    401,
		})
		return
	}
	unitName := data.Name

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	// 判断该类型是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM unit WHERE name=?", unitName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of unit for this unit name",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The unit already exists",
			"code":    401,
		})
		return
	}

	// 获取最近注册的单位的 unid
	var lastUnid string
	err = tx.QueryRow("SELECT unid FROM unit ORDER BY add_time DESC LIMIT 1").Scan(&lastUnid)
	// 如果没有单位，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastUnid = "un0000"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get last unid",
			"detail": err.Error(),
			"code":   503,
		})
		return
	}
	lastUnid = lastUnid[2:]
	// 增加最近注册的单位的 unid
	nextUnid, err := strconv.Atoi(lastUnid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot convert ITid to integer",
			"detail": err.Error(),
			"code":   504,
		})
		return
	}
	nextUnid++
	newUnid := fmt.Sprintf("un%04d", nextUnid) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	//增加仓库
	_, err = tx.Exec("INSERT INTO unit(unid, name, add_time) VALUES(?, ?, ?)", newUnid, unitName, addTime)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the unit",
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
		"message": "Unit added successfully",
		"code":    201,
	})
}
