package department

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	depName := data.Name

	db := myDb.GetMyDbConnection()

	// 判断该部门是否已经存在
	var dep model.Department
	err := db.Where("name = ?", depName).First(&dep).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The department name already exists",
			"code":    402,
		})
		return
	}
	newDid := "d" + utils.GenerateUuid(8) // 转换为 8 位字符串

	dep = model.Department{
		Did:  newDid,
		Name: depName,
	}

	err = db.Create(&dep).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert new department",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Department added successfully",
		"code":    201,
	})
}
