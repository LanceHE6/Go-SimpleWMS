package department

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
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

	err := db.Model(&dep).Where("did = ?", did).Updates(dep).Error
	if err != nil {
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
