package unit

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
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
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	unitName := data.Name

	db := my_db.GetMyDbConnection()

	// 判断该单位是否已经存在
	var unit model.Unit
	err := db.Where("name = ?", unitName).First(&unit).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "Unit already exists", nil))
		return
	}
	newUnid := "un" + utils.GenerateUuid(8) // 转换为 8 位字符串

	unit = model.Unit{
		Unid: newUnid,
		Name: unitName,
	}

	err = db.Create(&unit).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to add unit", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Unit added successfully", nil))
}
