package staff

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateRequest struct {
	Sid    string `json:"sid" form:"sid" binding:"required"`
	Name   string `json:"name" form:"name"`
	Phone  string `json:"phone" form:"phone"`
	DeptId string `json:"dept_id" form:"dept_id"`
}

func UpdateStaff(context *gin.Context) {
	var data updateRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Sid is required",
			"code":    401,
		})
		return
	}
	sid := data.Sid
	name := data.Name
	phone := data.Phone
	deptId := data.DeptId

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	if name == "" && phone == "" && deptId == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "At least one of name, dept_id and phone is required",
			"code":    402,
		})
		return
	}
	// 拼接sql语句
	updateSql := "UPDATE staff SET "
	if name != "" {
		updateSql += "name='" + name + "',"
	}
	if phone != "" {
		updateSql += "phone='" + phone + "',"
	}
	if deptId != "" {
		updateSql += "department=" + deptId + ","
	}
	updateSql = updateSql[:len(updateSql)-1] // 去掉最后一个逗号
	updateSql += " WHERE sid='" + sid + "'"
	_, err = tx.Exec(updateSql)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update staff",
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
		"message": "Staff updated successfully",
		"code":    201,
	})
}
