package warehouse

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteWarehouseRequest struct {
	Wid string `json:"wid" form:"name" binding:"required"`
}

func DeleteWarehouse(context *gin.Context) {
	var data deleteWarehouseRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "wid is required"})
		return
	}
	wid := data.Wid

	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 删除仓库
	_, err := tx.Exec("DELETE FROM warehouse WHERE wid=?", wid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete the warehouse"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit the transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Warehouse deleted successfully"})
}
