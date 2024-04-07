package department

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteDepartmentRequest struct {
	Did string `json:"did" form:"did" binding:"required"`
}

func DeleteDepartment(context *gin.Context) {
	var data deleteDepartmentRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	did := data.Did
	db := myDb.GetMyDbConnection()

	// 删除用户
	err := db.Delete(&model.Department{}, "did=?", did).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete department",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Department deleted successfully",
		"code":    201,
	})
}
