package user

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUser(context *gin.Context) {
	uid := context.PostForm("uid")
	password := context.PostForm("password")
	nickName := context.PostForm("nick_name")
	permission := context.PostForm("permission")

	if uid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "UID is required"})
		return
	}

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	if password == "" && nickName == "" && permission == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "At least one of password, nick_name and permission is required"})
		return
	}
	// 拼接sql语句
	updateSql := "UPDATE user SET "
	if password != "" {
		updateSql += "password='" + password + "',"
	}
	if nickName != "" {
		updateSql += "nick_name='" + nickName + "',"
	}
	if permission != "" {
		updateSql += "permission=" + permission + ","
	}
	updateSql = updateSql[:len(updateSql)-1] // 去掉最后一个逗号
	updateSql += " WHERE uid='" + uid + "'"
	_, err = tx.Exec(updateSql)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update user"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}
