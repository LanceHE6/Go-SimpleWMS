package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type updateInvRequest struct {
	Iid          string `json:"iid" form:"iid" binding:"required"`
	Date         string `json:"date" form:"date"`                     // 单据日期
	Number       string `json:"number" form:"number"`                 // 单号
	Department   string `json:"department" form:"department"`         // 单据所属部门
	GoodsList    string `json:"goods_list" form:"goods_list"`         // 单据包含的货品
	InventoryTpe string `json:"inventory_type" form:"inventory_type"` // 出入库类型
	Warehouse    string `json:"warehouse" form:"warehouse"`           // 所属仓库
	Operator     string `json:"operator" form:"operator"`             // 经办人
	Comment      string `json:"comment" form:"comment"`               // 备注
	Manufacturer string `json:"manufacturer" form:"manufacturer"`     // 制造商
}

func UpdateInv(context *gin.Context) {
	var data updateInvRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The format of the goods_list is incorrect",
			"code":    402,
			"detail":  err.Error(),
		})
		return
	}
	inventoryTpe := data.InventoryTpe
	warehouse := data.Warehouse
	operator := data.Operator
	comment := data.Comment
	manufacturer := data.Manufacturer

	parsedDate, _ := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
	// 构造更新数据
	var inv = map[string]interface{}{
		"date":           parsedDate,
		"number":         number,
		"department":     department,
		"goods_list":     goodsList,
		"inventory_type": inventoryTpe,
		"warehouse":      warehouse,
		"operator":       operator,
		"comment":        comment,
		"manufacturer":   manufacturer,
	}

	db := myDb.GetMyDbConnection()
	// 开始事务
	tx := db.Begin()
	if tx.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to start transaction",
			"code":    501,
			"detail":  tx.Error.Error(),
		})
		return
	}

	// 获取原来的出入库单
	var oldInv model.Inventory
	if err := tx.Model(&model.Inventory{}).Where("iid = ?", iid).First(&oldInv).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get original inventory record",
			"code":    502,
			"detail":  err.Error(),
		})
		return
	}

	// 获取出入库类型
	itid := oldInv.InventoryType
	var inventoryType model.InventoryType
	if err := tx.Model(&model.InventoryType{}).Where("itid = ?", itid).First(&inventoryType).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get inventory type",
			"code":    503,
			"detail":  err.Error(),
		})
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
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Contains invalid Stock records",
				"code":    403,
			})
			return
		}
		// 根据出入库类型回退数量
		if inventoryType.Type == 1 {
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
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to update stock",
				"code":    504,
				"detail":  err.Error(),
			})
			return
		}
	}

	// 更新出入库单
	if err := tx.Model(&model.Inventory{}).Where("iid = ?", iid).Updates(inv).Error; err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":  "The Number already exists",
				"detail": err.Error(),
				"code":   404,
			})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to update inventory record",
				"code":    505,
				"detail":  err.Error(),
			})
			return
		}
	}

	// 获取新的出入库单
	var newInv model.Inventory
	if err := tx.Model(&model.Inventory{}).Where("iid = ?", iid).First(&newInv).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get updated inventory record",
			"code":    506,
			"detail":  err.Error(),
		})
		return
	}

	// 获取出入库类型
	itid = newInv.InventoryType
	var newInventoryType model.InventoryType
	if err := tx.Model(&model.InventoryType{}).Where("itid = ?", itid).First(&newInventoryType).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get updated inventory type",
			"code":    507,
			"detail":  err.Error(),
		})
		return
	}

	// 更新货品数量
	if UpdateStocks(newInv.GoodsList, newInv.Warehouse, newInventoryType, context, tx) == 0 {
		// 提交事务
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to commit transaction",
				"code":    508,
				"detail":  err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Inventory updated successfully",
			"code":    201,
		})
	} else {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update stock quantities",
			"code":    509,
		})
	}
}
