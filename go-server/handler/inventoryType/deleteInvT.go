package inventoryType

import (
	"Go_simpleWMS/utils"
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
			"message": "ITid is required",
			"code":    401,
		})
		return
	}
	itid := data.ITid

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	// 删除仓库
	_, err = tx.Exec("DELETE FROM inventory_type WHERE itid=?", itid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the inventory type",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
			"code":   503,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Inventory type deleted successfully",
		"code":    201,
	})
}
