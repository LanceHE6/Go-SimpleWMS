package goods

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type updateGoodsRequest struct {
	Gid          string  `json:"gid" form:"gid" binding:"required"`
	Name         string  `json:"name" form:"name"`
	Model        string  `json:"model" form:"model"`
	GoodsCode    string  `json:"goods_code" form:"goods_code"`
	GoodsType    string  `json:"goods_type" form:"goods_type"`
	Manufacturer string  `json:"manufacturer" form:"manufacturer"`
	Unit         string  `json:"unit" form:"unit"`
	UnitPrice    float64 `json:"unit_price" form:"unit_price"`
}

func UpdateGoods(context *gin.Context) {
	var data updateGoodsRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	Gid := data.Gid
	GName := data.Name
	GModel := data.Model
	GCode := data.GoodsCode
	GType := data.GoodsType
	GManufacturer := data.Manufacturer
	GUnit := data.Unit
	GUnitPrice := data.UnitPrice

	var goods = utils.CreateUpdateData("name", GName,
		"model", GModel,
		"goods_code", GCode,
		"goods_type", GType,
		"manufacturer", GManufacturer,
		"unit", GUnit,
		"unit_price", GUnitPrice)

	db := my_db.GetMyDbConnection()

	err := db.Model(&model.Goods{}).Where("gid=?", Gid).Updates(goods).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			context.JSON(http.StatusOK, response.Response(403, "The goods already exist", nil))
			return
		}
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Update failed", err.Error()))
		return
	}
	context.JSON(http.StatusOK, response.Response(200, "Update successfully", nil))
}
