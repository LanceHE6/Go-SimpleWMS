package department

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteDepartmentRequest struct {
	Did string `json:"did" form:"did" binding:"required"`
}

func DeleteDepartment(context *gin.Context) {
	var data deleteDepartmentRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	did := data.Did
	db := myDb.GetMyDbConnection()

	// 删除用户
	err := db.Delete(&model.Department{}, "did=?", did).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to delete department", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Successfully deleted department", nil))
}
