package department

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListDepartment(context *gin.Context) {
	db := myDb.GetMyDbConnection()

	var departments []model.Department
	err := db.Select("*").Find(&departments).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Can not get the list of departments",
			"detail": err.Error(),
			"code":   201,
		})
	}
	// 封装返回列表
	var res []model.Department
	for _, department := range departments {
		res = append(res, department)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Get departments list successfully",
		"rows":    res,
		"code":    201,
	})
}
