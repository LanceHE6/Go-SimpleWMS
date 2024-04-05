package inventoryType

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type addInventoryTypeRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	TypeCode string `json:"type_code" form:"type_code"`
}

func AddInventoryType(context *gin.Context) {
	var data addInventoryTypeRequest
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

	// 判断该类型是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM inventory_type WHERE name=?", typeName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of inventory type for this type name",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The type name already exists",
			"code":    401,
		})
		return
	}

	newITid := "it" + utils.GenerateUuid(8) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	// 增加仓库
	_, err = tx.Exec("INSERT INTO inventory_type(itid, name, add_time, type_code) VALUES(?, ?, ?, ?)", newITid, typeName, addTime, typeCode)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the inventory type",
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
		"message": "Inventory type added successfully",
		"code":    201,
	})
}
