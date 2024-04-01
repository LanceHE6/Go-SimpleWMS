package warehouse

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListWarehouse(context *gin.Context) {
	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	rows, err := tx.Query("SELECT wid, name, add_time, comment, manager, status FROM warehouse")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of warehouses",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot close the list of warehouses",
				"detail": err.Error(),
				"code":   503,
			})
		}
	}(rows)

	var warehouses []gin.H
	for rows.Next() {
		var wid, name, addTime string

		//sql.NullString是一个结构体，它有两个字段：String和Valid。如果SQL查询结果中的值为NULL，Valid字段会被设置为false，否则为true
		var comment, manager sql.NullString
		var status int

		err = rows.Scan(&wid, &name, &addTime, &comment, &manager, &status)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot scan the list of warehouses",
				"detail": err.Error(),
				"code":   504,
			})
			return
		}
		var managerStr string
		if manager.Valid {
			managerStr = manager.String
		} else {
			managerStr = ""
		}
		var commentStr string
		if comment.Valid {
			commentStr = comment.String
		} else {
			commentStr = ""
		}

		warehouse := gin.H{
			"wid":      wid,
			"name":     name,
			"add_time": addTime,
			"comment":  commentStr,
			"manager":  managerStr,
			"status":   status,
		}
		warehouses = append(warehouses, warehouse)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get warehouse list successfully",
		"rows":    warehouses,
		"code":    201,
	})
}
