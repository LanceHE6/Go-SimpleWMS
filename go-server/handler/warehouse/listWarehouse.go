package warehouse

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListWarehouse(context *gin.Context) {
	db := myDb.GetMyDbConnection()
	var warehouses []model.Warehouse

	err := db.Select("*").Find(&warehouses).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Get warehouse list failed", err.Error()))
		return
	}

	var warehousesRes []model.Warehouse
	for _, warehouse := range warehouses {
		warehousesRes = append(warehousesRes, warehouse)
	}
	context.JSON(http.StatusOK, response.Response(200, "Get warehouse list successfully", gin.H{
		"rows": warehousesRes,
	}))
}
