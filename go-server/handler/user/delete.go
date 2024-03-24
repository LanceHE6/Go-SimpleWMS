package user

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(context *gin.Context) {
	uid := context.PostForm("uid")
	if uid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "UID is required"})
		return
	}
	tx, err := utils.GetDbConnection()
	// 开始一个新的事务
	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}
	// 删除用户
	_, err = tx.Exec("DELETE FROM user WHERE uid=?", uid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete user"})
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
