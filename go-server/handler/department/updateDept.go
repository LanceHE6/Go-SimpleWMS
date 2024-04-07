package department

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	did := data.Did
	depName := data.Name

	dep := model.Department{
		Name: depName,
	}

	db := myDb.GetMyDbConnection()
	// 判断该部门是否已存在
	err := db.Model(&model.Department{}).Where("did=?", did).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The department does not exist",
			"code":    403,
		})
		return
	}

	err = db.Model(&dep).Where("did = ?", did).Updates(dep).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":  "The name is already exists",
				"detail": err.Error(),
				"code":   402,
			})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update department",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Department updated successfully",
		"code":    201,
	})
}
