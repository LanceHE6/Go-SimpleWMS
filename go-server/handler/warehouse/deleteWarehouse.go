package warehouse

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteWarehouseRequest struct {
	Wid string `json:"wid" form:"name" binding:"required"`
}

func DeleteWarehouse(context *gin.Context) {
	var data deleteWarehouseRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	wid := data.Wid

	db := myDb.GetMyDbConnection()

	// 删除仓库
	var delData model.Warehouse
	db.Model(&model.Warehouse{}).Where("wid = ?", wid).First(&delData)
	err := db.Delete(&delData).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the warehouse",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Warehouse deleted successfully",
		"code":    201,
	})
}
