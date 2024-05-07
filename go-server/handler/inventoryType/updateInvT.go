package inventoryType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type updateInventoryTypeRequest struct {
	ITid     string `json:"itid" form:"itid" binding:"required"`
	Name     string `json:"name" form:"name"`
	TypeCode string `json:"type_code" form:"type_code"`
}

func UpdateInventoryType(context *gin.Context) {
	var data updateInventoryTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	ITid := data.ITid
	ITName := data.Name
	typeCode := data.TypeCode

	//if ITName == "" && typeCode == "" {
	//	context.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Name or type_code is required",
	//		"code":    402,
	//	})
	//	return
	//}

	db := myDb.GetMyDbConnection()

	var updateData = map[string]interface{}{
		"name":      ITName,
		"type_code": typeCode,
	}

	// 判断该类型是否已存在
	err := db.Model(&model.InventoryType{}).Where("itid=?", ITid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The inventory type does not exist",
			"code":    403,
		})
		return
	}
	err = db.Model(&model.InventoryType{}).Where("itid=?", ITid).Updates(updateData).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of inventory type for this itid",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Inventory type updated successfully",
		"code":    201,
	})
}
