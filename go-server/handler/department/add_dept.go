package department

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

type addDepartmentRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func AddDepartment(context *gin.Context) {
	var data addDepartmentRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	depName := data.Name

	db := my_db.GetMyDbConnection()

	// 判断该部门是否已经存在
	var dep model.Department
	err := db.Where("name = ?", depName).First(&dep).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "The department name already exists", nil))
		return
	}
	newDid := "d" + utils.GenerateUuid(8) // 转换为 8 位字符串

	dep = model.Department{
		Did:  newDid,
		Name: depName,
	}

	err = db.Create(&dep).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Cannot insert new department", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Successfully added new department", nil))
}
