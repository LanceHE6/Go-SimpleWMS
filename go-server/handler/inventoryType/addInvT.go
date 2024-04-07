package inventoryType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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

	db := myDb.GetMyDbConnection()
	// 判断该类型是否已存在
	var invt model.InventoryType
	err := db.Where(&invt, "name=?", typeName).Error

	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The type name already exists",
			"code":    401,
		})
		return
	}

	newITid := "it" + utils.GenerateUuid(8) // 转换为 8 位字符串

	invt = model.InventoryType{
		Name:     typeName,
		Itid:     newITid,
		TypeCode: typeCode,
	}
	// 增加仓库
	err = db.Create(&invt).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the inventory type",
			"detail": err.Error(),
			"code":   505,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Inventory type added successfully",
		"code":    201,
	})
}
