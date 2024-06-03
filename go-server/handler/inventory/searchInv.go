package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func SearchInv(context *gin.Context) {
	var invs []model.Inventory
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	goods := context.Query("goods")
	number := context.Query("number")
	warehouse := context.Query("warehouse")
	manufacturer := context.Query("manufacturer")
	amount, _ := strconv.Atoi(context.DefaultQuery("amount", "0"))
	inventoryType := context.Query("inventory_type")
	operator := context.Query("operator")
	comment := context.Query("comment")
	keyword := context.Query("keyword")

	query := myDb.GetMyDbConnection()
	// 计算偏移量
	offset := (page - 1) * limit

	if goods != "" {
		query = query.Where("goods = ?", goods)
	}
	if number != "" {
		query = query.Where("number = ?", number)
	}
	if warehouse != "" {
		query = query.Where("warehouse = ?", warehouse)
	}
	if manufacturer != "" {
		query = query.Where("manufacturer = ?", manufacturer)
	}
	if amount != 0 {
		query = query.Where("amount = ?", amount)
	}
	if inventoryType != "" {
		query = query.Where("inventory_type = ?", inventoryType)
	}
	if operator != "" {
		query = query.Where("operator = ?", operator)
	}
	if comment != "" {
		query = query.Where("comment = ?", comment)
	}
	if keyword != "" {
		query = query.Where("goods LIKE ? OR number LIKE ? OR warehouse LIKE ? OR manufacturer LIKE ? OR amount LIKE ? OR inventory_type LIKE ? OR operator LIKE ? OR comment LIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%")
	}

	// 获取总记录数
	var total int64
	query.Model(&model.Inventory{}).Count(&total)

	// 计算总页数
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	// 设置分页参数
	query = query.Offset(offset).Limit(limit)

	result := query.Offset(offset).Limit(limit).Find(&invs)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
			"code":  401,
		})
		return
	}
	if len(invs) == 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":        202,
			"message":     "No data",
			"page":        page,
			"page_size":   limit,
			"total":       total,
			"total_pages": totalPages,
			"rows":        invs,
		})
		return
	}

	var invsRes []gin.H
	for _, g := range invs {
		goodsMeta := gin.H{
			"created_at":     g.CreatedAt,
			"update_at":      g.UpdatedAt,
			"iid":            g.Iid,
			"number":         g.Number,
			"goods":          g.Goods,
			"inventory_type": g.InventoryType,
			"warehouse":      g.Warehouse,
			"manufacturer":   g.Manufacturer,
			"amount":         g.Amount,
			"operator":       g.Operator,
			"comment":        g.Comment,
		}
		invsRes = append(invsRes, goodsMeta)
	}

	context.JSON(http.StatusOK, gin.H{
		"code":        201,
		"message":     "Query successfully",
		"page":        page,
		"page_size":   limit,
		"total":       total,
		"total_pages": totalPages,
		"keyword":     keyword,
		"rows":        invsRes,
	})

}
