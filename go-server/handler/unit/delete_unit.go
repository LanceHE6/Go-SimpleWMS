package unit

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteUnitRequest struct {
	Unid string `json:"unid" form:"unid" binding:"required"`
}

func DeleteUnit(context *gin.Context) {
	var data deleteUnitRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	unid := data.Unid

	db := my_db.GetMyDbConnection()

	err := db.Delete(&model.Unit{}, "unid=?", unid).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to delete unit", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Unit deleted successfully", nil))
}
