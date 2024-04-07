package staff

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type addStaffRequest struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Phone  string `json:"phone" form:"phone"`
	DeptId string `json:"dept_id" form:"dept_id" binding:"required"`
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
	staffName := data.Name
	staffDeptId := data.DeptId
	phone := data.Phone

	db := myDb.GetMyDbConnection()

	var staff model.Staff

	// 判断该员工是否已存在
	err := db.Where("name=?", staffName).First(&staff).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The staff already exists",
			"code":    402,
		})
		return
	}

	newSid := "s" + utils.GenerateUuid(8) // 转换为 8 位字符串

	var dep model.Department
	err = db.Model(&model.Department{}).Where("did=?", staffDeptId).First(&dep).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The staff's department does not exist",
			"code":    403,
		})
		return
	}

	// 增加员工
	staff = model.Staff{
		Name:  staffName,
		Sid:   newSid,
		Phone: phone,
		Did:   staffDeptId,
	}
	err = db.Create(&staff).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the staff",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Staff added successfully",
		"code":    201,
	})
}
