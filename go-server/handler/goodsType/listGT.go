package goodsType

import (
	"Go_simpleWMS/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListGoodsType(context *gin.Context) {
	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	rows, err := tx.Query("SELECT gtid, name, type_code, add_time FROM goods_type")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the list of goods type"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot close the list of goods type"})
		}
	}(rows)

	var gts []gin.H
	for rows.Next() {
		var gtid, name, addTime string
		var typeCode sql.NullString

		err = rows.Scan(&gtid, &name, &typeCode, &addTime)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot scan the list of goods type" + err.Error()})
			return
		}
		var typeCodeStr string
		if typeCode.Valid {
			typeCodeStr = typeCode.String
		} else {
			typeCodeStr = ""
		}

		user := gin.H{
			"gtid":      gtid,
			"name":      name,
			"type_code": typeCodeStr,
			"addTime":   addTime,
		}
		gts = append(gts, user)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get user list successfully",
		"rows":    gts,
	})
}
