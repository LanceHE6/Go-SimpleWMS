package goods_type

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

type updateGoodsTypeRequest struct {
	GTid     string `json:"gtid" form:"gtid" binding:"required"`
	Name     string `json:"name" form:"name"`
	TypeCode string `json:"type_code" form:"type_code"`
}

func UpdateGoodsType(context *gin.Context) {
	var data updateGoodsTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	gTid := data.GTid
	gTName := data.Name
	typeCode := data.TypeCode

	//if gTName == "" && typeCode == "" {
	//	context.JSON(http.StatusBadRequest, gin.H{
	//		"message": "name or type_code is required",
	//		"code":    402,
	//	})
	//	return
	//}

	db := my_db.GetMyDbConnection()

	// 判断该类型是否已存在
	var gt model.GoodsType
	err := db.Model(&model.GoodsType{}).Where("gtid=?", gTid).First(&gt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "Goods type not found", nil))
		return
	}

	var updateData = map[string]interface{}{
		"name":      gTName,
		"type_code": typeCode,
	}

	err = db.Model(&model.GoodsType{}).Where("gtid=?", gTid).Updates(updateData).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			context.JSON(http.StatusOK, response.Response(403, "Goods type already exists", nil))
			return
		}
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Update failed", err.Error()))
		return
	}
	context.JSON(http.StatusOK, response.Response(200, "Update success", nil))
}
