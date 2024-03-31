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
	for _, userInfo := range userList {
		// 循环注册用户
		status, returnInfo := DoRegister(userInfo)
		if status != http.StatusOK {
			returnInfo["process"] = fmt.Sprintf("%d/%d", count, num)
			context.JSON(status, returnInfo)
			return
		}
		count++
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    201,
		"message": "All users have completed registration",
	})
}
