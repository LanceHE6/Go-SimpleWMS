package staff

import (
	"Go_simpleWMS/utils/response"
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
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
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
			context.JSON(http.StatusOK, response.Response(203, "All staffs failed to register", nil))
			return
		} else {
			context.JSON(http.StatusOK, response.Response(202, fmt.Sprintf("Some staffs failed to register; %d/%d", count, num), gin.H{
				"detail": errs,
			}))
			return
		}
	}
	context.JSON(http.StatusOK, response.Response(201, "All staffs registered successfully", nil))
}
