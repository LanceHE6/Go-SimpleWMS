package user

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUsers(context *gin.Context) {
	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	rows, err := tx.Query("SELECT uid, account, permission, register_time, phone, nick_name FROM user")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the list of users"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot close the list of users"})
		}
	}(rows)

	var users []gin.H
	for rows.Next() {
		var uid, account, registerTime, nickName string
		var permission int
		var phone sql.NullString

		err = rows.Scan(&uid, &account, &permission, &registerTime, &phone, &nickName)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot scan the list of users" + err.Error()})
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
			"nickName":      nickName,
		}
		users = append(users, user)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get user list successfully",
		"rows":    users,
	})
}
