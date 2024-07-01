package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteRequest struct {
	Uid string `json:"uid" form:"uid" binding:"required"`
}

func DeleteUser(context *gin.Context) {
	var data deleteRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	uid := data.Uid

	myClaims, err := utils.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, response.Response(101, "Unauthorized", nil))
		return
	}
	targetUid := myClaims.Uid
	if targetUid == uid {
		context.JSON(http.StatusOK, response.Response(102, "You cannot delete yourself", nil))
		return
	}
	db := my_db.GetMyDbConnection()

	// 删除用户
	err = db.Delete(&model.User{}, "uid=?", uid).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Internal Server Error", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Success", nil))
}
