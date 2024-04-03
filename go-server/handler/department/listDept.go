package department

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListDepartment(context *gin.Context) {
	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	rows, err := tx.Query("SELECT * FROM department")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of departments",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot close the list of departments",
				"detail": err.Error(),
				"code":   503,
			})
		}
	}(rows)

	var departments []gin.H
	for rows.Next() {
		var did, name, addTime string

		err = rows.Scan(&did, &name, &addTime)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot scan the list of department",
				"detail": err.Error(),
				"code":   504,
			})
			return
		}
		department := gin.H{
			"did":      did,
			"name":     name,
			"add_time": addTime,
		}
		departments = append(departments, department)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get departments list successfully",
		"rows":    departments,
		"code":    201,
	})
}
