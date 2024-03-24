package warehouse

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListWarehouse(context *gin.Context) {
	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	rows, err := tx.Query("SELECT wid, name, add_time, comment FROM warehouse")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the list of warehouses"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot close the list of warehouses"})
		}
	}(rows)

	var warehouses []gin.H
	for rows.Next() {
		var wid, name, addTime string

		//sql.NullString是一个结构体，它有两个字段：String和Valid。如果SQL查询结果中的值为NULL，Valid字段会被设置为false，否则为true
		var comment sql.NullString

		err = rows.Scan(&wid, &name, &addTime, &comment)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot scan the list of warehouses" + err.Error()})
			return
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
		}
		warehouses = append(warehouses, warehouse)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get warehouse list successfully",
		"rows":    warehouses,
	})
}
