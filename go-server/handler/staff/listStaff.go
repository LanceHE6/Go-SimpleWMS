package staff

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListStaff(context *gin.Context) {
	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	rows, err := tx.Query("SELECT * FROM staff")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of staffs",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot close the list of staffs",
				"detail": err.Error(),
				"code":   503,
			})
		}
	}(rows)

	var staffs []gin.H
	for rows.Next() {
		var sid, name, addTime, deptId string
		var phone sql.NullString

		err = rows.Scan(&sid, &name, &phone, &deptId, &addTime)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot scan the list of staffs",
				"detail": err.Error(),
				"code":   504,
			})
			return
		}
		var phoneStr string
		if phone.Valid {
			phoneStr = phone.String
		} else {
			phoneStr = ""
		}
		department := gin.H{
			"sid":        sid,
			"name":       name,
			"add_time":   addTime,
			"department": deptId,
			"phone":      phoneStr,
		}
		staffs = append(staffs, department)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get staffs list successfully",
		"rows":    staffs,
		"code":    201,
	})
}
