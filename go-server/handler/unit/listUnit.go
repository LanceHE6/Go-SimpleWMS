package unit

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUnit(context *gin.Context) {
	db := myDb.GetMyDbConnection()

	var units []model.Unit

	err := db.Select("*").Find(&units).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Can not get the list of units",
			"detail": err.Error(),
			"code":   201,
		})
	}
	// 封装返回列表
	var res []model.Unit
	for _, unit := range units {
		res = append(res, unit)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Get units list successfully",
		"rows":    res,
		"code":    201,
	})
}
