package warehouse

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListWarehouse(context *gin.Context) {
	db := myDb.GetMyDbConnection()
	var warehouses []model.Warehouse

	err := db.Select("*").Find(&warehouses).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of warehouses",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	var warehousesRes []model.Warehouse
	for _, warehouse := range warehouses {
		warehousesRes = append(warehousesRes, warehouse)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get warehouse list successfully",
		"rows":    warehousesRes,
		"code":    201,
	})
}
