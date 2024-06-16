package goodsType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListGoodsType(context *gin.Context) {
	db := myDb.GetMyDbConnection()
	var gts []model.GoodsType

	err := db.Select("*").Find(&gts).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Get goods type list failed", err.Error()))
		return
	}

	var gtsRes []model.GoodsType
	for _, gt := range gts {
		gtsRes = append(gtsRes, gt)
	}
	context.JSON(http.StatusOK, response.Response(200, "Get goods type list successfully", gin.H{
		"rows": gtsRes,
	}))
}
