package unit

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteUnitRequest struct {
	Unid string `json:"unid" form:"unid" binding:"required"`
}

func DeleteUnit(context *gin.Context) {
	var data deleteUnitRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	unid := data.Unid

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	// 删除单位
	_, err = tx.Exec("DELETE FROM unit WHERE unid=?", unid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the unit",
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
		"message": "Unit deleted successfully",
		"code":    201,
	})
}
