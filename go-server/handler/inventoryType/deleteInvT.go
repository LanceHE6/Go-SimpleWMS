package inventoryType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteInventoryTypeRequest struct {
	ITid string `json:"itid" form:"itid" binding:"required"`
}

func DeleteInventoryType(context *gin.Context) {
	var data deleteInventoryTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	itid := data.ITid

	db := myDb.GetMyDbConnection()

	// 删除仓库
	err := db.Delete(&model.InventoryType{}, "itid=?", itid).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the inventory type",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Inventory type deleted successfully",
		"code":    201,
	})
}
