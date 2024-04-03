package unit

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUnit(context *gin.Context) {
	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	rows, err := tx.Query("SELECT unid, name FROM unit")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of unit",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot close the list of unit",
				"detail": err.Error(),
				"code":   503,
			})
		}
	}(rows)

	var unitArray []gin.H
	for rows.Next() {
		var unid, name string

		err = rows.Scan(&unid, &name)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot scan the list of unit",
				"detail": err.Error(),
				"code":   504,
			})
			return
		}

		unit := gin.H{
			"unid": unid,
			"name": name,
		}
		unitArray = append(unitArray, unit)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get inventory type list successfully",
		"rows":    unitArray,
		"code":    201,
	})
}
