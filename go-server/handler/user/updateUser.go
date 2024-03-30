package user

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type updateRequest struct {
	Uid        string `json:"uid" form:"uid" binding:"required"`
	Password   string `json:"password" form:"password"`
	Nickname   string `json:"nickname" form:"nickname"`
	Permission int    `json:"permission" form:"permission"`
	Phone      string `json:"phone" form:"phone"`
}

func UpdateUser(context *gin.Context) {
	var data updateRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "UID is required",
			"code":    401,
		})
		return
	}
	uid := data.Uid
	password := data.Password
	nickName := data.Nickname
	permission := data.Permission
	phone := data.Phone

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	if password == "" && nickName == "" && permission == 0 && phone == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "At least one of password, nickname, permission and phone is required",
			"code":    402,
		})
		return
	}
	// 拼接sql语句
	updateSql := "UPDATE user SET "
	if password != "" {
		updateSql += "password='" + password + "',"
	}
	if nickName != "" {
		updateSql += "nickname='" + nickName + "',"
	}
	if permission != 0 {
		updateSql += "permission=" + strconv.Itoa(permission) + ","
	}
	if phone != "" {
		updateSql += "phone='" + phone + "',"
	}
	updateSql = updateSql[:len(updateSql)-1] // 去掉最后一个逗号
	updateSql += " WHERE uid='" + uid + "'"
	_, err = tx.Exec(updateSql)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update user",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit transaction",
			"detail": err.Error(),
			"code":   503,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"code":    201,
	})
}
