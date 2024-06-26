package goods

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteGoodsRequest struct {
	Gid string `json:"gid" form:"gid" binding:"required"`
}

func DeleteGoods(context *gin.Context) {
	var data deleteGoodsRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	gid := data.Gid

	db := my_db.GetMyDbConnection()

	// 删除
	var goods model.Goods
	db.Model(&model.Goods{}).Where("gid=?", gid).First(&goods)
	err := db.Delete(&goods).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to delete goods", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Successfully deleted goods", nil))
}
