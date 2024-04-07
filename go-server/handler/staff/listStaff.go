package staff

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListStaff(context *gin.Context) {
	db := myDb.GetMyDbConnection()
	var staffs []model.Staff

	err := db.Select([]string{"sid", "name", "phone", "created_at"}).Find(&staffs).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of staffs",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}

	var staffsRes []gin.H
	for _, staff := range staffs {

		staffMeta := gin.H{
			"sid":        staff.Sid,
			"name":       staff.Name,
			"created_at": staff.CreatedAt,
			"department": staff.Did,
			"phone":      staff.Phone,
		}
		staffsRes = append(staffsRes, staffMeta)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get staffs list successfully",
		"rows":    staffsRes,
		"code":    201,
	})
}
