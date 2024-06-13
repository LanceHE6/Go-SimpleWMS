package staff

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"fmt"
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
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
		return http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("The staff's %s department does not exist", staff.Name),
			"code":    403,
		}
	}
	if err := db.Create(&staff).Error; err != nil {
		return http.StatusInternalServerError, gin.H{
			"error":  fmt.Sprintf("Cannot insert new staff %s", staff.Name),
			"detail": err.Error(),
			"code":   505,
		}
	}

	return http.StatusOK, gin.H{
		"message": "Staff registered successfully",
		"sid":     newSid,
		"code":    201,
	}
}
