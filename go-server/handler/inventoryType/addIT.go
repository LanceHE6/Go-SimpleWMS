package inventoryType

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

type addInventoryTypeRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	TypeCode string `json:"type_code" form:"type_code"`
}

func AddInventoryType(context *gin.Context) {
	var data addInventoryTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Type name is required"})
		return
	}
	typeName := data.Name
	typeCode := data.TypeCode

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 判断该类型是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM inventory_type WHERE name=?", typeName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the number of inventory type for this type name"})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{"message": "The type name already exists"})
		return
	}

	// 获取最近注册的货品类型的 gtid
	var lastITid string
	err = tx.QueryRow("SELECT itid FROM inventory_type ORDER BY add_time DESC LIMIT 1").Scan(&lastITid)
	// 如果没有用户，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastITid = "it0000"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get last ITid"})
		return
	}
	lastITid = lastITid[2:]
	// 增加最近注册的用户的 uid
	nextITid, err := strconv.Atoi(lastITid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot convert ITid to integer"})
		return
	}
	nextITid++
	newITid := fmt.Sprintf("it%04d", nextITid) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	// 增加仓库
	_, err = tx.Exec("INSERT INTO inventory_type(itid, name, add_time, type_code) VALUES(?, ?, ?, ?)", newITid, typeName, addTime, typeCode)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot insert the inventory type"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit the transaction"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Goods type added successfully"})
}
