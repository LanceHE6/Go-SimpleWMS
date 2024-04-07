package unit

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type addUnitRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func AddUnit(context *gin.Context) {
	var data addUnitRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	unitName := data.Name

	db := myDb.GetMyDbConnection()

	// 判断该单位是否已经存在
	var unit model.Unit
	err := db.Where("name = ?", unitName).First(&unit).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The unit already exists",
			"code":    402,
		})
		return
	}
	newUnid := "un" + utils.GenerateUuid(8) // 转换为 8 位字符串

	unit = model.Unit{
		Unid: newUnid,
		Name: unitName,
	}

	err = db.Create(&unit).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert new unit",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Unit added successfully",
		"code":    201,
	})
}
