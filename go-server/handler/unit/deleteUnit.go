package unit

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
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

	db := myDb.GetMyDbConnection()

	err := db.Delete(&model.Unit{}, "unid=?", unid).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the unit",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Unit deleted successfully",
		"code":    201,
	})
}
