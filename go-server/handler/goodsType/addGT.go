package goodsType

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type addGoodsTypeRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	TypeCode string `json:"type_code" form:"type_code"`
}

func AddGoodsType(context *gin.Context) {
	var data addGoodsTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	typeName := data.Name
	typeCode := data.TypeCode

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
	err = tx.QueryRow("SELECT count(name) FROM goods_type WHERE name=?", typeName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of goods type for this type name",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The type name already exists",
			"code":    402,
		})
		return
	}

	newGTid := "gt" + utils.GenerateUuid(8) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	// 增加仓库
	_, err = tx.Exec("INSERT INTO goods_type(gtid, name, add_time, type_code) VALUES(?, ?, ?, ?)", newGTid, typeName, addTime, typeCode)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the goods type",
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
		"message": "Goods type added successfully",
		"code":    201,
	})
}
