package goods

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type addGoodsRequest struct {
	Name         string `json:"name" form:"name" binding:"required"`
	Model        string `json:"model" form:"model"`
	GoodsCode    string `json:"goods_code" form:"goods_code"`
	GoodsType    string `json:"goods_type" form:"goods_type" binding:"required"`
	Warehouse    string `json:"warehouse" form:"warehouse" binding:"required"`
	Manufacturer string `json:"manufacturer" form:"manufacturer"`
	Unit         string `json:"unit" form:"unit" binding:"required"`
	Quantity     int    `json:"quantity" form:"quantity"`
}

func AddGoods(context *gin.Context) {
	var data addGoodsRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	GName := data.Name
	GModel := data.Model
	GCode := data.GoodsCode
	GType := data.GoodsType
	GWarehouse := data.Warehouse
	GManufacturer := data.Manufacturer
	GUnit := data.Unit
	GQuantity := data.Quantity

	db := myDb.GetMyDbConnection()

	// 判断该仓库是否已存在
	var goods model.Goods
	if GCode != "" {
		err := db.Model(&model.Goods{}).Where("goods_code=?", GCode).First(&goods).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusForbidden, gin.H{
				"message": "The goods with this code already exists",
				"code":    402,
			})
			return
		}
	}

	newGid := "g" + utils.GenerateUuid(8) // 转换为 8 位字符串

	// 增加仓库
	goods = model.Goods{
		Gid:          newGid,
		Name:         GName,
		Model:        GModel,
		GoodsCode:    GCode,
		GoodsType:    GType,
		Warehouse:    GWarehouse,
		Manufacturer: GManufacturer,
		Unit:         GUnit,
		Quantity:     GQuantity,
	}
	err := db.Create(&goods).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the goods",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Goods added successfully",
		"code":    201,
	})
}
