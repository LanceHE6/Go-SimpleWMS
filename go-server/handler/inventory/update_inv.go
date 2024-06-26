package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/handler/stock"
	"Go_simpleWMS/utils/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type updateInvRequest struct {
	Iid           string `json:"iid" form:"iid" binding:"required"`
	Date          string `json:"date" form:"date"`                     // 单据日期
	Number        string `json:"number" form:"number"`                 // 单号
	Department    string `json:"department" form:"department"`         // 单据所属部门
	GoodsList     string `json:"goods_list" form:"goods_list"`         // 单据包含的货品
	InventoryType string `json:"inventory_type" form:"inventory_type"` // 出入库类型
	Warehouse     string `json:"warehouse" form:"warehouse"`           // 所属仓库
	Operator      string `json:"operator" form:"operator"`             // 经办人
	Comment       string `json:"comment" form:"comment"`               // 备注
	Manufacturer  string `json:"manufacturer" form:"manufacturer"`     // 制造商
}

func UpdateInv(context *gin.Context) {
	var data updateInvRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}

	// 获取请求参数
	iid := data.Iid
	date := data.Date
	number := data.Number
	department := data.Department
	var goodsList model.GoodsList
	err := json.Unmarshal([]byte(data.GoodsList), &goodsList)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.Response(402, "Invalid goods list", nil))
		return
	}
	iTid := data.InventoryType
	warehouse := data.Warehouse
	operator := data.Operator
	comment := data.Comment
	manufacturer := data.Manufacturer

	parsedDate, _ := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)

	db := my_db.GetMyDbConnection()
	// 开始事务
	tx := db.Begin()
	if tx.Error != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to start transaction", tx.Error.Error()))
		return
	}

	// 获取原来的出入库单
	var oldInv model.Inventory
	if err := tx.Model(&model.Inventory{}).Where("iid = ?", iid).First(&oldInv).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(502, "Failed to get inventory", err.Error()))
		return
	}

	// 获取出入库类型
	itid := oldInv.InventoryType
	var oldInventoryType model.InventoryType
	if err := tx.Model(&model.InventoryType{}).Where("itid = ?", itid).First(&oldInventoryType).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(503, "Failed to get inventory type", err.Error()))
		return
	}

	// 根据原来的出入库单回退对应的货品数量
	for _, goodsInfo := range oldInv.GoodsList {
		var oldStock = model.Stock{
			Goods:     goodsInfo.Goods,
			Warehouse: warehouse,
		}
		notExist := tx.Model(&model.Stock{}).Where("goods = ? AND warehouse = ?", oldStock.Goods, oldInv.Warehouse).First(&oldStock).RecordNotFound()
		if notExist {
			tx.Rollback()
			context.JSON(http.StatusOK, response.Response(403, "Failed to get stock", nil))
			return
		}
		// 根据出入库类型回退数量
		if oldInventoryType.Type == 1 {
			oldStock.Quantity -= goodsInfo.Amount
		} else {
			oldStock.Quantity += goodsInfo.Amount
		}
		var updateData = map[string]interface{}{
			"quantity": oldStock.Quantity,
		}
		fmt.Println("old:", oldStock)
		if err := tx.Model(&model.Stock{}).Where("goods = ? AND warehouse = ?", oldStock.Goods, oldStock.Warehouse).Updates(updateData).Error; err != nil {
			tx.Rollback()
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(504, "Failed to update stock", err.Error()))
			return
		}
	}

	// 获取更新后的出入库类型
	var newInventoryType model.InventoryType
	if err := tx.Model(&model.InventoryType{}).Where("itid = ?", iTid).First(&newInventoryType).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(505, "Failed to get inventory type", err.Error()))
		return
	}
	// 构造更新后库存数据
	var oldGoodsList = oldInv.OldGoodsList
	var newGoodsList model.GoodsList
	for _, g := range oldGoodsList {
		var newGoodsOrder model.GoodsOrder
		newGoodsOrder.Goods = g.Goods
		oldStock := stock.GetGoodsListAmountByGoods(oldGoodsList, g.Goods)
		if newInventoryType.Type == 1 {
			newGoodsOrder.Amount = oldStock + g.Amount
		} else {
			newGoodsOrder.Amount = oldStock - g.Amount
		}
		newGoodsList = append(newGoodsList, newGoodsOrder)
	}
	// 如果新库存数据中有库存不足的情况：存在负数
	for _, g := range newGoodsList {
		if g.Amount < 0 {
			tx.Rollback()
			context.JSON(http.StatusBadRequest, response.Response(404, "Contains invalid Stock records", nil))
			return
		}
	}

	// 构造更新数据
	var inv = map[string]interface{}{
		"date":           parsedDate,
		"number":         number,
		"department":     department,
		"goods_list":     goodsList,
		"old_goods_list": oldGoodsList,
		"new_goods_list": newGoodsList,
		"inventory_type": iTid,
		"warehouse":      warehouse,
		"operator":       operator,
		"comment":        comment,
		"manufacturer":   manufacturer,
	}
	// 更新出入库单
	if err := tx.Model(&model.Inventory{}).Where("iid = ?", iid).Updates(inv).Error; err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			context.JSON(http.StatusOK, response.Response(405, "Duplicate inventory number", nil))
			return
		} else {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(506, "Failed to update inventory", err.Error()))
			return
		}
	}

	// 获取新的出入库单
	var newInv model.Inventory
	if err := tx.Model(&model.Inventory{}).Where("iid = ?", iid).First(&newInv).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(507, "Failed to get inventory", err.Error()))
		return
	}

	// 更新货品数量
	if stock.UpdateStocks(newInv.GoodsList, newInv.Warehouse, newInventoryType, context, tx) == 0 {
		// 提交事务
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(508, "Failed to commit transaction", err.Error()))
			return
		}

		context.JSON(http.StatusOK, response.Response(200, "Successfully updated inventory", nil))
	}
}
