package user

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUsers(context *gin.Context) {
	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   "501",
		})
		return
	}

	rows, err := tx.Query("SELECT uid, account, permission, register_time, phone, nickname FROM user")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of users",
			"detail": err.Error(),
			"code":   "502",
		})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot close the list of users",
				"detail": err.Error(),
				"code":   "503",
			})
		}
	}(rows)

	var users []gin.H
	for rows.Next() {
		var uid, account, registerTime, nickName string
		var permission int
		var phone sql.NullString

		err = rows.Scan(&uid, &account, &permission, &registerTime, &phone, &nickName)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot scan the list of users",
				"detail": err.Error(),
				"code":   "504",
			})
			return
		}
		var phoneStr string
		if phone.Valid {
			phoneStr = phone.String
		} else {
			phoneStr = ""
		}

		user := gin.H{
			"uid":           uid,
			"account":       account,
			"permission":    permission,
			"register_time": registerTime,
			"phone":         phoneStr,
			"nickname":      nickName,
		}
		users = append(users, user)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get user list successfully",
		"rows":    users,
		"code":    201,
	})
}
