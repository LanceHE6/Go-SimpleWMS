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
	err := db.Select("*").Where("is_deleted=?", 0).Find(&invTs).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of inventory type",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	var invTsRes []model.InventoryType
	for _, invT := range invTs {
		invTsRes = append(invTsRes, invT)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get inventory type list successfully",
		"rows":    invTsRes,
		"code":    201,
	})
}
