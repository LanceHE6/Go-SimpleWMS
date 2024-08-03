package inventory_type

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type addInventoryTypeRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	TypeCode string `json:"type_code" form:"type_code"`
	Type     int    `json:"type" form:"type" binding:"required"`
}

func AddInventoryType(context *gin.Context) {
	var data addInventoryTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	typeName := data.Name
	typeCode := data.TypeCode
	typeNum := data.Type

	db := my_db.GetMyDbConnection()
	// 判断该类型是否已存在
	var invt model.InventoryType
	notExist := db.Model(model.InventoryType{}).Where("name=?", typeName).First(&invt).RecordNotFound()

	if !notExist {
		context.JSON(http.StatusOK, response.Response(402, "The inventory type already exists", nil))
		return
	}

	newITid := "it" + utils.GenerateUUID(8) // 转换为 8 位字符串

	invt = model.InventoryType{
		Name:     typeName,
		Itid:     newITid,
		TypeCode: typeCode,
		Type:     typeNum,
	}
	// 添加
	err := db.Create(&invt).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Cannot create the inventory type", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Create inventory type successfully", nil))
}
