package unit

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateUnitRequest struct {
	Unid string `json:"unid" form:"unid" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

func UpdateUnit(context *gin.Context) {
	var data updateUnitRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	unid := data.Unid
	name := data.Name

	db := myDb.GetMyDbConnection()
	// 判断该单位是否存在
	var unit model.Unit
	notFound := db.Model(&model.Unit{}).Where("unid=?", unid).First(&unit).RecordNotFound()

	if notFound {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The unit does not exist",
			"code":    403,
		})
		return
	}

	var updateData = map[string]interface{}{
		"name": name,
	}
	err := db.Model(&model.Unit{}).Where("unid=?", unid).Updates(updateData).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update unit",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Unit updated successfully",
		"code":    201,
	})
}
