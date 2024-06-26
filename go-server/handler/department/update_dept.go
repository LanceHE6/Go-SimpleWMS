package department

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

type updateDeptRequest struct {
	Did  string `json:"did" form:"did" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

func UpdateDepartment(context *gin.Context) {
	var data updateDeptRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	did := data.Did
	depName := data.Name

	dep := model.Department{
		Name: depName,
	}

	db := my_db.GetMyDbConnection()
	// 判断该部门是否已存在
	err := db.Model(&model.Department{}).Where("did=?", did).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "Department not found", nil))
		return
	}

	err = db.Model(&dep).Where("did = ?", did).Updates(dep).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			context.JSON(http.StatusOK, response.Response(403, "Department already exists", nil))
			return
		}
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to update department", err.Error()))
		return
	}
	context.JSON(http.StatusOK, response.Response(200, "Department updated successfully", nil))
}
