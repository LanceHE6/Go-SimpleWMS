package unit

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateUnitRequest struct {
	Unid string `json:"unid" form:"unid" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

func UpdateUnit(context *gin.Context) {
	var data updateUnitRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	unid := data.Unid
	name := data.Name

	db := myDb.GetMyDbConnection()
	// 判断该单位是否存在
	var unit model.Unit
	notFound := db.Model(&model.Unit{}).Where("unid=?", unid).First(&unit).RecordNotFound()

	if notFound {
		context.JSON(http.StatusOK, response.Response(402, "Unit not found", nil))
		return
	}

	var updateData = map[string]interface{}{
		"name": name,
	}
	err := db.Model(&model.Unit{}).Where("unid=?", unid).Updates(updateData).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Update unit failed", err.Error()))
		return
	}
	context.JSON(http.StatusOK, response.Response(200, "Update unit success", nil))
}
