package department

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListDepartment(context *gin.Context) {
	db := myDb.GetMyDbConnection()

	var departments []model.Department
	err := db.Select("*").Find(&departments).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Get departments list failed", err.Error()))
	}
	// 封装返回列表
	var res []model.Department
	for _, department := range departments {
		res = append(res, department)
	}

	context.JSON(http.StatusOK, response.Response(200, "Get departments list successfully", gin.H{
		"rows": res,
	}))
}
