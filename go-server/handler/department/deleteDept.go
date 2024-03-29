package department

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteDepartmentRequest struct {
	Did string `json:"did" form:"did" binding:"required"`
}

func DeleteDepartment(context *gin.Context) {
	var data deleteDepartmentRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "did is required"})
		return
	}
	did := data.Did

	tx, err := utils.GetDbConnection()
	// 开始一个新的事务
	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}
	// 删除部门
	_, err = tx.Exec("DELETE FROM department WHERE did=?", did)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete department"})
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Department deleted successfully",
	})
}