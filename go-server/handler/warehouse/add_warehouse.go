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

type addWarehouseRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Comment string `json:"comment" form:"comment"`
	Manager string `json:"manager" form:"manager" binding:"required"`
	Status  int    `json:"status" form:"status"`
}

func AddWarehouse(context *gin.Context) {
	var data addWarehouseRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	warehouseName := data.Name
	comment := data.Comment
	manager := data.Manager
	status := data.Status

	db := my_db.GetMyDbConnection()

	// 判断该仓库是否已存在
	var warehouse model.Warehouse
	err := db.Model(&model.Warehouse{}).Where("name=?", warehouseName).First(&warehouse).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "Warehouse already exists", nil))
		return
	}

	newWid := "w" + utils.GenerateUUID(8)
	warehouse = model.Warehouse{
		Wid:     newWid,
		Name:    warehouseName,
		Comment: comment,
		Manager: manager,
		Status:  status,
	}
	err = db.Create(&warehouse).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to add warehouse", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Successfully added warehouse", nil))
}
