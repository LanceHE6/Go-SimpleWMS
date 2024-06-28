package warehouse

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type updateWarehouseRequest struct {
	Wid     string `json:"wid" form:"wid" binding:"required"`
	Name    string `json:"name" form:"name"`
	Comment string `json:"comment" form:"comment"`
	Manager string `json:"manager" form:"manager"`
	Status  int    `json:"status" form:"status"`
}

func UpdateWarehouse(context *gin.Context) {
	var data updateWarehouseRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	wid := data.Wid
	warehouseName := data.Name
	comment := data.Comment
	manager := data.Manager
	status := data.Status

	db := my_db.GetMyDbConnection()

	// 判断该仓库是否已存在
	err := db.Model(&model.Warehouse{}).Where("wid=?", wid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "Warehouse not found", nil))
		return
	}

	// 使用map封装更新数据，能更新所有数据，包括零值字段
	var updateData = utils.CreateUpdateData("name", warehouseName, "comment", comment, "manager", manager, "status", status)

	err = db.Model(&model.Warehouse{}).Where("wid=?", wid).Updates(updateData).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to update warehouse", err.Error()))
		return
	}
	context.JSON(http.StatusOK, response.Response(200, "Successfully updated warehouse", nil))
}
