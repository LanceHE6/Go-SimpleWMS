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

	err := db.Select([]string{"wid", "name", "created_at", "comment", "manager", "status"}).Find(&warehouses).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of warehouses",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	var warehousesRes []gin.H
	for _, warehouse := range warehouses {

		warehouseRes := gin.H{
			"wid":        warehouse.Wid,
			"name":       warehouse.Name,
			"created_at": warehouse.CreatedAt,
			"comment":    warehouse.Comment,
			"manager":    warehouse.Manager,
			"status":     warehouse.Status,
		}
		warehousesRes = append(warehousesRes, warehouseRes)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get warehouse list successfully",
		"rows":    warehousesRes,
		"code":    201,
	})
}
