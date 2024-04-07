package warehouse

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	warehouseName := data.Name
	comment := data.Comment
	manager := data.Manager
	status := data.Status

	db := myDb.GetMyDbConnection()

	// 判断该仓库是否已存在
	var warehouse model.Warehouse
	err := db.Model(&model.Warehouse{}).Where("name=?", warehouseName).First(&warehouse).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusConflict, gin.H{
			"message": "The warehouse already exists",
			"code":    402,
		})
		return
	}

	newWid := "w" + utils.GenerateUuid(8)
	warehouse = model.Warehouse{
		Wid:     newWid,
		Name:    warehouseName,
		Comment: comment,
		Manager: manager,
		Status:  status,
	}
	err = db.Create(&warehouse).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the warehouse",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Warehouse added successfully",
		"code":    201,
	})
}
