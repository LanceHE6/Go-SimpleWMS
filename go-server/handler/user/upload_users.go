package user

import (
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type uploadUsersRequest struct {
	List []registerRequest `json:"list" form:"list" binding:"required"`
}

func UploadUsers(context *gin.Context) {
	var data uploadUsersRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}

	userList := data.List
	num := len(userList)
	var errs []gin.H
	for _, userInfo := range userList {
		// 循环注册用户
		status, returnInfo := DoRegister(userInfo)
		if status != http.StatusOK {
			errs = append(errs, returnInfo)
			continue
		}
	}
	if len(errs) > 0 {
		if len(errs) == num {
			context.JSON(http.StatusOK, response.Response(203, "All users failed to register", gin.H{
				"detail": errs,
			}))

			return
		} else {
			context.JSON(http.StatusOK, response.Response(202, "Some users failed to register", gin.H{
				"detail": errs,
			}))
			return
		}
	}
	context.JSON(http.StatusOK, response.Response(201, "All users registered successfully", nil))
}
