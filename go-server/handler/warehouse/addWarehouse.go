package warehouse

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

type addWarehouseRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Comment string `json:"comment" form:"comment"`
	Manager string `json:"manager" form:"manager" binding:"required"`
	Status  string `json:"status" form:"status"`
}

func AddWarehouse(context *gin.Context) {
	var data addWarehouseRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Warehouse name and manager are required",
			"code":    401,
		})
		return
	}
	warehouseName := data.Name
	comment := data.Comment
	manager := data.Manager
	status := data.Status

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	// 判断该仓库是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM warehouse WHERE name=?", warehouseName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of warehouses for this warehouse_name",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusConflict, gin.H{
			"message": "The warehouse already exists",
			"code":    409,
		})
		return
	}

	// 获取最近注册的仓库的 wid
	var lastWid string
	err = tx.QueryRow("SELECT wid FROM warehouse ORDER BY add_time DESC LIMIT 1").Scan(&lastWid)
	// 如果没有用户，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastWid = "wh0000"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get last wid",
			"detail": err.Error(),
			"code":   503,
		})
		return
	}
	lastWid = lastWid[2:]
	// 增加最近注册的仓库的 wid
	nextWid, err := strconv.Atoi(lastWid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot convert wid to integer",
			"detail": err.Error(),
			"code":   504,
		})
		return
	}
	nextWid++
	newWid := fmt.Sprintf("wh%04d", nextWid) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	// 增加仓库
	if comment == "" {
		if status == "" {
			_, err = tx.Exec("INSERT INTO warehouse(wid, name, add_time, manager) VALUES(?, ?, ?, ?)", newWid, warehouseName, addTime, manager)
		} else {
			_, err = tx.Exec("INSERT INTO warehouse(wid, name, add_time, manager, status) VALUES(?, ?, ?, ?, ?)", newWid, warehouseName, addTime, manager, status)
		}
	} else {
		if status == "" {
			_, err = tx.Exec("INSERT INTO warehouse(wid, name, add_time, manager, comment) VALUES(?, ?, ?, ?, ?)", newWid, warehouseName, addTime, manager, comment)
		} else {
			_, err = tx.Exec("INSERT INTO warehouse(wid, name, add_time, manager, comment, status) VALUES(?, ?, ?, ?, ?, ?)", newWid, warehouseName, addTime, manager, comment, status)
		}
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the warehouse",
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

	context.JSON(http.StatusCreated, gin.H{
		"message": "Warehouse added successfully",
		"code":    201,
	})
}
