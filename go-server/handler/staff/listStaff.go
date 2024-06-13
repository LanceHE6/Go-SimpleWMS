package staff

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListStaff(context *gin.Context) {
	db := myDb.GetMyDbConnection()
	var staffs []model.Staff

	err := db.Select("*").Find(&staffs).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Get staffs list failed", err.Error()))
		return
	}

	var staffsRes []model.Staff
	for _, staff := range staffs {
		staffsRes = append(staffsRes, staff)
	}
	context.JSON(http.StatusOK, response.Response(200, "Get staffs list successfully", gin.H{
		"rows": staffsRes,
	}))
}
