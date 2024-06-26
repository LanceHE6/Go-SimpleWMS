package inventory_type

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type updateInventoryTypeRequest struct {
	ITid     string `json:"itid" form:"itid" binding:"required"`
	Name     string `json:"name" form:"name"`
	TypeCode string `json:"type_code" form:"type_code"`
	Type     int    `json:"type" form:"type"`
}

func UpdateInventoryType(context *gin.Context) {
	var data updateInventoryTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	ITid := data.ITid
	ITName := data.Name
	typeCode := data.TypeCode
	typeNum := data.Type

	//if ITName == "" && typeCode == "" {
	//	context.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Name or type_code is required",
	//		"code":    402,
	//	})
	//	return
	//}

	db := my_db.GetMyDbConnection()

	var updateData = map[string]interface{}{
		"name":      ITName,
		"type_code": typeCode,
		"type":      typeNum,
	}

	// 判断该类型是否已存在
	err := db.Model(&model.InventoryType{}).Where("itid=?", ITid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "Inventory type not found", nil))
		return
	}
	err = db.Model(&model.InventoryType{}).Where("itid=?", ITid).Updates(updateData).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to update inventory type", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Inventory type updated successfully", nil))
}
