package staff

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type addStaffRequest struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Phone  string `json:"phone" form:"phone"`
	DeptId string `json:"department" form:"department" binding:"required"`
}

func AddStaff(context *gin.Context) {
	var data addStaffRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	// 执行注册逻辑
	status, returnData := DoAddStaff(data)

	context.JSON(status, returnData)
}

// DoAddStaff 执行注册逻辑函数
func DoAddStaff(staffData addStaffRequest) (int, gin.H) {
	name := staffData.Name
	phone := staffData.Phone
	deptId := staffData.DeptId

	db := myDb.GetMyDbConnection()

	// 插入新用户
	newSid := "s" + utils.GenerateUuid(8)
	staff := model.Staff{
		Sid:        newSid,
		Name:       name,
		Phone:      phone,
		Department: deptId,
	}
	var dep model.Department
	err := db.Model(&model.Department{}).Where("did=?", staff.Department).First(&dep).Error
	if err != nil {
		return http.StatusOK, response.Response(402, "Department not found", nil)
	}
	if err := db.Create(&staff).Error; err != nil {
		return http.StatusInternalServerError, response.ErrorResponse(501, "Failed to register staff", err.Error())
	}

	return http.StatusOK, response.Response(200, "Register success", nil)
}
