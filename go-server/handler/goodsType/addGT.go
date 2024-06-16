package goodsType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type addGoodsTypeRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	TypeCode string `json:"type_code" form:"type_code"`
}

func AddGoodsType(context *gin.Context) {
	var data addGoodsTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	typeName := data.Name
	typeCode := data.TypeCode

	db := myDb.GetMyDbConnection()

	// 判断该仓库是否已存在
	var gt model.GoodsType
	err := db.Model(&model.GoodsType{}).Where("name=?", typeName).First(&gt).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "The goods type already exists", nil))
		return
	}

	newGTid := "gt" + utils.GenerateUuid(8) // 转换为 8 位字符串

	// 增加仓库
	gt = model.GoodsType{
		Gtid:     newGTid,
		Name:     typeName,
		TypeCode: typeCode,
	}
	err = db.Create(&gt).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to add goods type", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Successfully added goods type", nil))
}
