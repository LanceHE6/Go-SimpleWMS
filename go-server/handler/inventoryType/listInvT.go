package inventoryType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListInventoryType(context *gin.Context) {
	db := myDb.GetMyDbConnection()
	var invTs []model.InventoryType
	err := db.Select([]string{"itid", "name", "type_code", "type", "created_at"}).Find(&invTs).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of inventory type",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	var invTsRes []gin.H
	for _, invT := range invTs {

		invTRes := gin.H{
			"itid":       invT.Itid,
			"name":       invT.Name,
			"type_code":  invT.TypeCode,
			"created_at": invT.CreatedAt,
			"type":       invT.Type,
		}
		invTsRes = append(invTsRes, invTRes)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get inventory type list successfully",
		"rows":    invTsRes,
		"code":    201,
	})
}
