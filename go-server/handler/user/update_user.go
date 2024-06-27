package user

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type updateRequest struct {
	Uid         string `json:"uid" form:"uid" binding:"required"`
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password"`
	Nickname    string `json:"nickname" form:"nickname"`
	Permission  int    `json:"permission" form:"permission"`
	Phone       string `json:"phone" form:"phone"`
}

func UpdateUser(context *gin.Context) {
	var data updateRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	targetUid := data.Uid
	oldPassword := data.OldPassword
	newPassword := data.NewPassword
	nickname := data.Nickname
	permission := data.Permission
	phone := data.Phone

	db := my_db.GetMyDbConnection()

	var oldUser model.User
	// 判断该用户是否已存在
	err := db.Model(&model.User{}).Where("uid=?", targetUid).First(&oldUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusOK, response.Response(402, "User not found", nil))
		return
	}

	var updateData = make(map[string]interface{})
	if nickname != "" {
		updateData["nickname"] = nickname
	}
	if permission != 0 {
		// 权限只有超管能改
		if oldUser.Permission != 3 {
			context.JSON(http.StatusOK, response.Response(300, "Permission denied", nil))
			return
		}
		updateData["permission"] = permission
	}
	if phone != "" {
		updateData["phone"] = phone
	}
	myUid, myPermission, _, _ := utils.GetUserInfoByContext(context)
	// 如果要改密码:密码不为空
	if newPassword != "" {
		// 如果权限为管理员以上且改密码的用户不是自己
		if myPermission >= 2 {
			if targetUid != myUid {
				updateData["password"] = newPassword
			} else {
				// 改自己密码
				if oldUser.Password == oldPassword {
					updateData["password"] = newPassword
				} else {
					// 旧密码不正确
					context.JSON(http.StatusOK, response.Response(403, "Old password is incorrect", nil))
					return
				}
			}
		} else {
			// 权限为普通用户 只能更改自己的密码
			if targetUid == myUid {
				if oldUser.Password == oldPassword {
					updateData["password"] = newPassword
				} else {
					// 旧密码不正确
					context.JSON(http.StatusOK, response.Response(403, "Old password is incorrect", nil))
					return
				}
			} else {
				// 权限不足
				context.JSON(http.StatusOK, response.Response(300, "Permission denied", nil))
				return
			}
		}

	}

	err = db.Model(&model.User{}).Where("uid = ?", targetUid).Updates(updateData).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Update user failed", err.Error()))
		return
	}

	context.JSON(http.StatusOK, response.Response(200, "Update user successfully", nil))
}
