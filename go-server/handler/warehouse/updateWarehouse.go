package warehouse

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	wid := data.Wid
	warehouseName := data.Name
	comment := data.Comment
	manager := data.Manager
	status := data.Status

	if warehouseName == "" && comment == "" && manager == "" && status == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "One of name, comment, manager and status is required",
			"code":    402,
		})
		return
	}

	db := myDb.GetMyDbConnection()

	// 判断该仓库是否已存在
	err := db.Model(&model.Warehouse{}).Where("wid=?", wid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The department does not exist",
			"code":    403,
		})
		return
	}

	var warehouse = model.Warehouse{
		Wid:     wid,
		Name:    warehouseName,
		Comment: comment,
		Manager: manager,
		Status:  status,
	}
	err = db.Model(&model.Warehouse{}).Where("wid=?", warehouse.Wid).Updates(&warehouse).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update the warehouse",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Warehouse updated successfully",
		"code":    201,
	})
}
