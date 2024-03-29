package user

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteRequest struct {
	Uid string `json:"uid" form:"uid" binding:"required"`
}

func DeleteUser(context *gin.Context) {
	var data deleteRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "UID is required"})
		return
	}
	uid := data.Uid

	targetUid, err := utils.GetUidByContext(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	if targetUid == uid {
		context.JSON(http.StatusForbidden, gin.H{"message": "Invalid target uid"})
		return
	}
	tx, err := utils.GetDbConnection()
	// 开始一个新的事务
	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
		})
		return
	}
	// 删除用户
	_, err = tx.Exec("DELETE FROM user WHERE uid=?", uid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete user",
			"detail": err.Error(),
		})
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit transaction",
			"detail": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
