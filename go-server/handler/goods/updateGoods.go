package goods

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type updateGoodsRequest struct {
	Gid          string `json:"gid" form:"gid" binding:"required"`
	Name         string `json:"name" form:"name"`
	Model        string `json:"model" form:"model"`
	GoodsCode    string `json:"goods_code" form:"goods_code"`
	GoodsType    string `json:"goods_type" form:"goods_type"`
	Warehouse    string `json:"warehouse" form:"warehouse"`
	Manufacturer string `json:"manufacturer" form:"manufacturer"`
	Unit         string `json:"unit" form:"unit"`
	Quantity     int    `json:"quantity" form:"quantity"`
}

func UpdateGoods(context *gin.Context) {
	var data updateGoodsRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	Gid := data.Gid
	GName := data.Name
	GModel := data.Model
	GCode := data.GoodsCode
	GType := data.GoodsType
	GWarehouse := data.Warehouse
	GManufacturer := data.Manufacturer
	GUnit := data.Unit
	GQuantity := data.Quantity

	//if GName == "" && GModel == "" && GCode == "" && GType == "" && GWarehouse == "" && GManufacturer == "" && GUnit == "" && GQuantity == 0 {
	//	context.JSON(http.StatusBadRequest, gin.H{
	//		"message": "No data to update",
	//		"code":    402,
	//	})
	//	return
	//}

	var goods = map[string]interface{}{
		"name":         GName,
		"model":        GModel,
		"goods_code":   GCode,
		"goods_type":   GType,
		"warehouse":    GWarehouse,
		"manufacturer": GManufacturer,
		"unit":         GUnit,
		"quantity":     GQuantity,
	}

	db := myDb.GetMyDbConnection()

	err := db.Model(&model.Goods{}).Where("gid=?", Gid).Updates(goods).Error
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
			"message": "Update goods failed",
			"code":    403,
			"detail":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Goods updated successfully",
		"code":    201,
	})
}
