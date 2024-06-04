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
	var total int64
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	goods := context.Query("goods")
	number := context.Query("number")
	warehouse := context.Query("warehouse")
	manufacturer := context.Query("manufacturer")
	amount, _ := strconv.Atoi(context.DefaultQuery("amount", "0"))
	inventoryType := context.Query("inventory_type")
	iType, _ := strconv.Atoi(context.DefaultQuery("type", "0"))
	operator := context.Query("operator")
	comment := context.Query("comment")
	keyword := context.Query("keyword")
	createdAt := context.Query("created_at")

	query := myDb.GetMyDbConnection().Table("inventories").Joins("left join inventory_types on inventories.inventory_type = inventory_types.itid")

	if goods != "" {
		query = query.Where("inventories.goods = ?", goods)
	}
	if number != "" {
		query = query.Where("inventories.number = ?", number)
	}
	if warehouse != "" {
		query = query.Where("inventories.warehouse = ?", warehouse)
	}
	if manufacturer != "" {
		query = query.Where("inventories.manufacturer = ?", manufacturer)
	}
	if amount != 0 {
		query = query.Where("inventories.amount = ?", amount)
	}
	if inventoryType != "" {
		query = query.Where("inventories.inventory_type = ?", inventoryType)
	}
	if iType != 0 {
		query = query.Where("inventory_types.type = ?", iType)
	}
	if operator != "" {
		query = query.Where("inventories.operator = ?", operator)
	}
	if comment != "" {
		query = query.Where("inventories.comment = ?", comment)
	}
	if createdAt != "" {
		query = query.Where("inventories.created_at = ?", createdAt)
	}
	if keyword != "" {
		query = query.Where("inventories.goods LIKE ? OR inventories.number LIKE ? OR inventories.warehouse LIKE ? OR inventories.manufacturer LIKE ? OR inventories.amount LIKE ? OR inventories.inventory_type LIKE ? OR inventories.operator LIKE ? OR inventories.comment LIKE ? OR inventories.created_at LIKE ?",
			"%"+keyword+"%",
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
	query.Count(&total)

	var totalPages = 0

	// page 为-1不分页
	if page != -1 {
		// 计算总页数
		totalPages = int(math.Ceil(float64(total) / float64(limit)))
		// 计算偏移量
		offset := (page - 1) * limit
		// 设置分页参数
		query = query.Offset(offset).Limit(limit)
	}

	// 执行查询
	result := query.Find(&invs)
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
	db := myDb.GetMyDbConnection()
	for _, g := range invs {
		var goods model.Goods
		db.Model(model.Goods{}).Where("gid = ?", g.Goods).First(&goods)

		goodsMeta := gin.H{
			"created_at": g.CreatedAt,
			"update_at":  g.UpdatedAt,
			"iid":        g.Iid,
			"number":     g.Number,
			"goods": gin.H{
				"created_at":   goods.CreatedAt,
				"updated_at":   goods.UpdatedAt,
				"gid":          goods.Gid,
				"goods_code":   goods.GoodsCode,
				"name":         goods.Name,
				"model":        goods.Model,
				"goods_type":   goods.GoodsType,
				"manufacturer": goods.Manufacturer,
				"unit":         goods.Unit,
				"image":        goods.Image,
				"quantity":     goods.Quantity,
				"unit_price":   goods.UnitPrice,
			},
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
