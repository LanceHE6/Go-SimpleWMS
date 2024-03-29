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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "wid is required",
			"code":    "401",
		})
		return
	}
	wid := data.Wid

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
	_, err = tx.Exec("DELETE FROM warehouse WHERE wid=?", wid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the warehouse",
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
		"message": "Warehouse deleted successfully",
		"code":    201,
	})
}
