package inventoryType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListInventoryType(context *gin.Context) {
	db := myDb.GetMyDbConnection()
	var invTs []model.InventoryType
	err := db.Select("*").Where("is_deleted=?", 0).Find(&invTs).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Response(501, "Get inventory type list failed", nil))
		return
	}

	var invTsRes []model.InventoryType
	for _, invT := range invTs {
		invTsRes = append(invTsRes, invT)
	}
	context.JSON(http.StatusOK, response.Response(200, "Get inventory type list successfully", gin.H{
		"rows": invTsRes,
	}))
}
