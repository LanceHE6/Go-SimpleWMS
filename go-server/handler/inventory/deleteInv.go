package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type delInvRequest struct {
	Iid string `json:"iid" form:"iid" binding:"required"`
}

func DeleteInv(context *gin.Context) {
	var data delInvRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	iid := data.Iid

	db := myDb.GetMyDbConnection()
	tx := db.Begin()

	var inv model.Inventory

	// 检查仓库是否存在
	notExist := tx.Model(model.Inventory{}).Where("iid=?", iid).First(&inv).RecordNotFound()
	if notExist {
		context.JSON(http.StatusOK, response.Response(402, "Inventory not exist", nil))
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
				context.JSON(http.StatusInternalServerError, response.Response(503, "Cannot update the stock", nil))
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
				context.JSON(http.StatusInternalServerError, response.Response(504, "Cannot update the stock", nil))
			}

		}

	}
	// 删除单据
	err := tx.Delete(model.Inventory{}, "iid=?", iid).Error
	if err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, response.Response(505, "Cannot delete the inventory", nil))
		return
	}
	err = tx.Commit().Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Response(506, "Cannot commit the transaction", nil))
		return
	}
	context.JSON(http.StatusOK, response.Response(200, "Success", nil))
}
