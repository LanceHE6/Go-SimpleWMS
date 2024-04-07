package goodsType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	gTid := data.GTid
	gTName := data.Name
	typeCode := data.TypeCode

	if gTName == "" && typeCode == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "name or type_code is required",
			"code":    402,
		})
		return
	}

	db := myDb.GetMyDbConnection()

	// 判断该类型是否已存在
	var gt model.GoodsType
	err := db.Model(&model.GoodsType{}).Where("gtid=?", gTid).First(&gt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The goods type does not exist",
			"code":    403,
		})
		return
	}

	// 更新仓库
	gt = model.GoodsType{
		Gtid:     gTid,
		Name:     gTName,
		TypeCode: typeCode,
	}

	err = db.Model(&model.GoodsType{}).Where("gtid=?", gt.Gtid).Updates(&gt).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":  "The name is already exists",
				"detail": err.Error(),
				"code":   404,
			})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update the goods type",
			"detail": err.Error(),
			"code":   504,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Goods type updated successfully",
		"code":    201,
	})
}
