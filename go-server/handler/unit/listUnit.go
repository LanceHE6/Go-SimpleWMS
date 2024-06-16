package unit

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUnit(context *gin.Context) {
	db := myDb.GetMyDbConnection()

	var units []model.Unit

	err := db.Select("*").Find(&units).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Get units list failed", err.Error()))
	}
	// 封装返回列表
	var res []model.Unit
	for _, unit := range units {
		res = append(res, unit)
	}

	context.JSON(http.StatusOK, response.Response(200, "Get units list successfully", gin.H{
		"rows": res,
	}))
}
