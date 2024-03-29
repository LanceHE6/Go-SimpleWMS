package staff

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteStaffRequest struct {
	Sid string `json:"sid" form:"sid" binding:"required"`
}

func DeleteStaff(context *gin.Context) {
	var data deleteStaffRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "SID is required"})
		return
	}
	sid := data.Sid

	tx, err := utils.GetDbConnection()
	// 开始一个新的事务
	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}
	// 删除用户
	_, err = tx.Exec("DELETE FROM staff WHERE sid=?", sid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete staff"})
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Staff deleted successfully",
	})
}
