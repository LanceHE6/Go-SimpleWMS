package staff

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type uploadUsersRequest struct {
	StaffList []addStaffRequest `json:"list" form:"list" binding:"required"`
}

func UploadStaffs(context *gin.Context) {
	var data uploadUsersRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}

	staffList := data.StaffList
	//fmt.Println(userList)
	num := len(staffList)
	count := 0
	var errs []gin.H
	for _, staffInfo := range staffList {
		// 循环注册用户
		status, returnInfo := DoAddStaff(staffInfo)
		if status != http.StatusOK {
			errs = append(errs, returnInfo)
			continue
		}
		count++
	}
	if len(errs) > 0 {
		if len(errs) == num {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("All staffs failed to register"),
				"code":    402,
				"detail":  errs,
			})
			return
		} else {
			context.JSON(217, gin.H{
				"message": fmt.Sprintf("Some staffs failed to register; %d/%d", count, num),
				"code":    403,
				"detail":  errs,
			})
			return
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    201,
		"message": "All staffs have completed registration",
	})
}
