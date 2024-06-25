package staff

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
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
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
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

	db := my_db.GetMyDbConnection()
	// 判断该员工是否已存在
	var staff model.Staff
	notFound := db.Model(&model.Staff{}).Where("sid=?", sid).First(&staff).RecordNotFound()
	if notFound {
		context.JSON(http.StatusOK, response.Response(402, "The staff does not exist", nil))
		return
	}

	var dep model.Department

	if deptId != "" {
		err := db.Model(&model.Department{}).Where("did=?", deptId).Find(&dep).Error
		if err != nil {
			context.JSON(http.StatusOK, response.Response(403, "The department does not exist", nil))
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
		context.JSON(http.StatusInternalServerError, response.Response(501, "Update Error", nil))
		return
	}
	context.JSON(http.StatusOK, response.Response(200, "Success", nil))
}
