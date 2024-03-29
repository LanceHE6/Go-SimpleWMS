package inventoryType

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListInventoryType(context *gin.Context) {
	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	rows, err := tx.Query("SELECT itid, name, type_code, add_time FROM inventory_type")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of inventory type",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot close the list of inventory type",
				"detail": err.Error(),
				"code":   503,
			})
		}
	}(rows)

	var gts []gin.H
	for rows.Next() {
		var itid, name, addTime string
		var typeCode sql.NullString

		err = rows.Scan(&itid, &name, &typeCode, &addTime)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot scan the list of inventory type",
				"detail": err.Error(),
				"code":   504,
			})
			return
		}
		var typeCodeStr string
		if typeCode.Valid {
			typeCodeStr = typeCode.String
		} else {
			typeCodeStr = ""
		}

		user := gin.H{
			"gtid":      itid,
			"name":      name,
			"type_code": typeCodeStr,
			"addTime":   addTime,
		}
		gts = append(gts, user)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get inventory type list successfully",
		"rows":    gts,
		"code":    201,
	})
}
