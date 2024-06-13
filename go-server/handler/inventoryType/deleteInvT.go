package inventoryType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteInventoryTypeRequest struct {
	ITid string `json:"itid" form:"itid" binding:"required"`
}

func DeleteInventoryType(context *gin.Context) {
	var data deleteInventoryTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	itid := data.ITid

	db := myDb.GetMyDbConnection()

	// 软删除
	err := db.Model(&model.InventoryType{}).Where("itid = ?", itid).Update("is_deleted", 1).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to delete inventory type", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Successfully deleted inventory type", nil))
}
