package staff

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	tx, err := utils.GetDbConnection()
	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	// 判断该仓库是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM staff WHERE name=?", staffName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of staffs for this staff name",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The staff already exists",
			"code":    402,
		})
		return
	}

	newSid := "s" + utils.GenerateUuid(8) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	// 增加仓库
	_, err = tx.Exec("INSERT INTO staff(sid, name, add_time, phone, department) VALUES(?, ?, ?, ?, ?)", newSid, staffName, addTime, phone, staffDeptId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the staff",
			"detail": err.Error(),
			"code":   505,
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
			"code":   506,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Staff added successfully",
		"code":    201,
	})
}
