package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type delInvRequest struct {
	Iid string `json:"iid" form:"iid" binding:"required"`
}

func DeleteInv(context *gin.Context) {
	var data delInvRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	iid := data.Iid

	db := myDb.GetMyDbConnection()
	tx := db.Begin()

	var inv model.Inventory

	// 检查仓库是否存在
	notExist := tx.Model(model.Inventory{}).Where("iid=?", iid).First(&inv).RecordNotFound()
	if notExist {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Inventory not found",
			"code":  402,
		})
		return
	}
	warehouse := inv.Warehouse
	goodsList := inv.GoodsList
	var invT model.InventoryType
	tx.Model(model.InventoryType{}).Where("itid=?", inv.InventoryType).First(&invT)
	// 根据出入库方向回退货品数量
	if invT.Type == 1 {
		for _, goodsOrder := range goodsList {
			var stock model.Stock
			tx.Model(model.Stock{}).Where("warehouse=? and goods=?", warehouse, goodsOrder.Goods).First(&stock)
			var updateData = map[string]interface{}{
				"quantity": stock.Quantity - goodsOrder.Amount,
			}

			err := tx.Model(model.Stock{}).Where("warehouse=? and goods=?", warehouse, goodsOrder.Goods).Updates(updateData).Error
			if err != nil {
				tx.Rollback()
				context.JSON(http.StatusInternalServerError, gin.H{
					"error":  "Cannot update the stock",
					"detail": err.Error(),
					"code":   502,
				})
				return
			}
		}
	} else {
		for _, goodsOrder := range goodsList {
			var stock model.Stock
			tx.Model(model.Stock{}).Where("warehouse=? and goods=?", warehouse, goodsOrder.Goods).First(&stock)
			var updateData = map[string]interface{}{
				"quantity": stock.Quantity + goodsOrder.Amount,
			}
			err := tx.Model(model.Stock{}).Where("warehouse=? and goods=?", warehouse, goodsOrder.Goods).Updates(updateData).Error
			if err != nil {
				tx.Rollback()
				context.JSON(http.StatusInternalServerError, gin.H{
					"error":  "Cannot update the stock",
					"detail": err.Error(),
					"code":   503,
				})
			}

		}

	}
	// 删除单据
	err := tx.Delete(model.Inventory{}, "iid=?", iid).Error
	if err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the inventory",
			"detail": err.Error(),
			"code":   504,
		})
		return
	}
	err = tx.Commit().Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
			"code":   505,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "inventory deleted successfully",
		"code":    201,
	})
}
