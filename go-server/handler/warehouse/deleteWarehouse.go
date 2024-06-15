package warehouse

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteWarehouseRequest struct {
	Wid string `json:"wid" form:"name" binding:"required"`
}

func DeleteWarehouse(context *gin.Context) {
	var data deleteWarehouseRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	wid := data.Wid

	db := myDb.GetMyDbConnection()

	// 删除仓库
	var delData model.Warehouse
	db.Model(&model.Warehouse{}).Where("wid = ?", wid).First(&delData)
	err := db.Delete(&delData).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to delete warehouse", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Successfully deleted warehouse", nil))
}
