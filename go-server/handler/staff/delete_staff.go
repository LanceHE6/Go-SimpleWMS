package staff

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteStaffRequest struct {
	Sid string `json:"sid" form:"sid" binding:"required"`
}

func DeleteStaff(context *gin.Context) {
	var data deleteStaffRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	sid := data.Sid

	db := my_db.GetMyDbConnection()

	err := db.Delete(&model.Staff{}, "sid=?", sid).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to delete staff", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Staff deleted successfully", nil))
}
