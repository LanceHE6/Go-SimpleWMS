package staff

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteStaffRequest struct {
	Sid string `json:"sid" form:"sid" binding:"required"`
}

func DeleteStaff(context *gin.Context) {
	var data deleteStaffRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	sid := data.Sid

	db := myDb.GetMyDbConnection()

	err := db.Delete(&model.Staff{}, "sid=?", sid).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete staff",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Staff deleted successfully",
		"code":    201,
	})
}
