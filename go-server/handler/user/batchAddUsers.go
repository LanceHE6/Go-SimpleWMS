package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerBatchRequest struct {
	UserList []registerRequest `json:"user_list" form:"user_list" binding:"required"`
}

func RegisterBatch(context *gin.Context) {
	var data registerBatchRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}

	userList := data.UserList
	//fmt.Println(userList)
	num := len(userList)
	count := 0
	var errs []gin.H
	for _, userInfo := range userList {
		// 循环注册用户
		status, returnInfo := DoRegister(userInfo)
		if status != http.StatusOK {
			errs = append(errs, returnInfo)
			continue
		}
		count++
	}
	if len(errs) > 0 {
		if len(errs) == num {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("All users failed to register"),
				"code":    402,
				"detail":  errs,
			})
			return
		} else {
			context.JSON(217, gin.H{
				"message": fmt.Sprintf("Some users failed to register; %d/%d", count, num),
				"code":    403,
				"detail":  errs,
			})
			return
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    201,
		"message": "All users have completed registration",
	})
}
