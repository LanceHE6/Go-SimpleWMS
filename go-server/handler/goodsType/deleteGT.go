package goodsType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteGoodsTypeRequest struct {
	GTid string `json:"gtid" form:"gtid" binding:"required"`
}

func DeleteGoodsType(context *gin.Context) {
	var data deleteGoodsTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	gtid := data.GTid

	db := myDb.GetMyDbConnection()

	// 删除仓库
	err := db.Delete(&model.GoodsType{}, "gtid=?", gtid).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to delete goods type", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Successfully deleted goods type", nil))
}
