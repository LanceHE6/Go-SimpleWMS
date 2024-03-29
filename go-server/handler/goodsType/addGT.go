package goodsType

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

type addGoodsTypeRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	TypeCode string `json:"type_code" form:"type_code"`
}

func AddGoodsType(context *gin.Context) {
	var data addGoodsTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Type name is required"})
		return
	}
	typeName := data.Name
	typeCode := data.TypeCode

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
		})
		return
	}

	// 判断该仓库是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM goods_type WHERE name=?", typeName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of goods type for this type name",
			"detail": err.Error(),
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{"message": "The type name already exists"})
		return
	}

	// 获取最近注册的货品类型的 gtid
	var lastGTid string
	err = tx.QueryRow("SELECT gtid FROM goods_type ORDER BY add_time DESC LIMIT 1").Scan(&lastGTid)
	// 如果没有用户，就从 1 开始
	if errors.Is(err, sql.ErrNoRows) {
		lastGTid = "gt0000"
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get last GTid",
			"detail": err.Error(),
		})
		return
	}
	lastGTid = lastGTid[2:]
	// 增加最近注册的用户的 uid
	nextGTid, err := strconv.Atoi(lastGTid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot convert GTid to integer",
			"detail": err.Error(),
		})
		return
	}
	nextGTid++
	newGTid := fmt.Sprintf("gt%04d", nextGTid) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	// 增加仓库
	_, err = tx.Exec("INSERT INTO goods_type(gtid, name, add_time, type_code) VALUES(?, ?, ?, ?)", newGTid, typeName, addTime, typeCode)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the goods type",
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

	context.JSON(http.StatusOK, gin.H{"message": "Goods type added successfully"})
}
