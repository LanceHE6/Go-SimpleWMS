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

func AddWarehouse(context *gin.Context) {
	warehouseName := context.PostForm("name")
	comment := context.PostForm("comment")

	if warehouseName == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Warehouse name is required"})
		return
	}

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 判断该仓库是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM warehouse WHERE name=?", warehouseName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the number of warehouses for this warehouse_name"})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{"message": "The warehouse already exists"})
		return
	}

	// 获取最近注册的仓库的 wid
	var lastWid string
	err = tx.QueryRow("SELECT wid FROM warehouse ORDER BY add_time DESC LIMIT 1").Scan(&lastWid)
	// 如果没有用户，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastWid = "000000"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get last wid"})
		return
	}

	// 增加最近注册的用户的 uid
	nextWid, err := strconv.Atoi(lastWid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot convert wid to integer"})
		return
	}
	nextWid++
	newWid := fmt.Sprintf("%06d", nextWid) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	// 增加仓库
	if comment == "" {
		_, err = tx.Exec("INSERT INTO warehouse(wid, name, add_time) VALUES(?, ?, ?)", newWid, warehouseName, addTime)
	} else {
		_, err = tx.Exec("INSERT INTO warehouse(wid, name, add_time, comment) VALUES(?, ?, ?, ?)", newWid, warehouseName, addTime, comment)
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot insert the warehouse"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit the transaction"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Warehouse added successfully"})
}

func DeleteWarehouse(context *gin.Context) {
	wid := context.PostForm("wid")
	if wid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "wid is required"})
		return
	}

	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 删除仓库
	_, err := tx.Exec("DELETE FROM warehouse WHERE wid=?", wid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete the warehouse"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit the transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Warehouse deleted successfully"})
}
