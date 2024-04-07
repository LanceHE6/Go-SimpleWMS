package goodsType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
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

	// 判断该仓库是否已存在
	var gt model.GoodsType
	err := db.Model(&model.GoodsType{}).Where("name=?", typeName).First(&gt).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The type name already exists",
			"code":    402,
		})
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
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the goods type",
			"detail": err.Error(),
			"code":   505,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Goods type added successfully",
		"code":    201,
	})
}
