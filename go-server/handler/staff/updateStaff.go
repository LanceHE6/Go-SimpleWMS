package staff

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
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
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	sid := data.Sid
	name := data.Name
	phone := data.Phone
	deptId := data.DeptId

	//if name == "" && phone == "" && deptId == "" {
	//	context.JSON(http.StatusBadRequest, gin.H{
	//		"message": "One of name, phone, dept_id and phone is required",
	//		"code":    402,
	//	})
	//	return
	//}

	db := myDb.GetMyDbConnection()
	// 判断该员工是否已存在
	var staff model.Staff
	notFound := db.Model(&model.Staff{}).Where("sid=?", sid).First(&staff).RecordNotFound()
	if notFound {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The staff does not exist",
			"code":    403,
		})
		return
	}

	var dep model.Department

	if deptId != "" {
		err := db.Model(&model.Department{}).Where("did=?", deptId).Find(&dep).Error
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "The staff's department does not exist",
				"code":    404,
			})
			return
		}
	}

	var updateData = map[string]interface{}{
		"name":       name,
		"phone":      phone,
		"department": deptId,
	}
	err := db.Model(&model.Staff{}).Where("sid=?", sid).Updates(updateData).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update staff",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Staff updated successfully",
		"code":    201,
	})
}
