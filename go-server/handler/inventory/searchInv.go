package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"time"
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
	date := context.Query("date")

	query := myDb.GetMyDbConnection().Table("inventories").Joins("left join inventory_types on inventories.inventory_type = inventory_types.itid")
	// 依据参数构造查询
	if goods != "" {
		query = query.Where(query.Where("JSON_CONTAINS(inventories.goods_list, JSON_OBJECT('Goods', ?))", goods))
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
		query = query.Where("JSON_CONTAINS(inventories.goods_list, JSON_OBJECT('Amount', ?))", amount)
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
		query = query.Where("inventories.comment = ? OR JSON_CONTAINS(inventories.goods_list, JSON_OBJECT('Comment', ?))", comment, comment)
	}
	if date != "" {
		// 将 created_at 转换为日期格式，并过滤出当天的记录
		parsedDate, err := time.Parse("2006-01-02", date)
		fmt.Println(parsedDate)
		if err == nil {
			startOfDay := parsedDate.Format("2006-01-02 00:00:00")
			endOfDay := parsedDate.Add(24 * time.Hour).Format("2006-01-02 15:04:05")
			query = query.Where("inventories.date BETWEEN ? AND ?", startOfDay, endOfDay)
		}
	}
	if keyword != "" {
		query = query.Where("inventories.goods_list LIKE ? OR inventories.number LIKE ? OR inventories.warehouse LIKE ? OR inventories.manufacturer LIKE ? OR inventories.inventory_type LIKE ? OR inventories.operator LIKE ? OR inventories.comment LIKE ? OR inventories.date LIKE ?",
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
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Internal Server Error", result.Error.Error()))
		return
	}

	// 查询结果数量为0
	if len(invs) == 0 {
		context.JSON(http.StatusOK, response.Response(202, "No data", gin.H{
			"page":        page,
			"page_size":   limit,
			"total":       total,
			"total_pages": totalPages,
			"rows":        invs,
		}))
		return
	}

	var invsRes []gin.H
	db := myDb.GetMyDbConnection()
	for _, g := range invs {

		var goodsList []gin.H
		// 构造货品清单返回体
		for _, goodsInfo := range g.GoodsList {
			var goods model.Goods
			db.Model(model.Goods{}).Where("gid = ?", goodsInfo.Goods).First(&goods)
			goodsList = append(goodsList, gin.H{
				"goods":   goods,
				"amount":  goodsInfo.Amount,
				"comment": goodsInfo.Comment,
			})
		}
		// 构造单个出入库单返回体
		goodsMeta := gin.H{
			"created_at":     g.CreatedAt,
			"update_at":      g.UpdatedAt,
			"iid":            g.Iid,
			"number":         g.Number,
			"goods_list":     goodsList,
			"old_goods_list": g.OldGoodsList,
			"new_goods_list": g.NewGoodsList,
			"warehouse":      g.Warehouse,
			"inventory_type": g.InventoryType,
			"manufacturer":   g.Manufacturer,
			"operator":       g.Operator,
			"comment":        g.Comment,
		}
		invsRes = append(invsRes, goodsMeta)
	}

	context.JSON(http.StatusOK, response.Response(201, "Query successfully", gin.H{
		"page":        page,
		"page_size":   limit,
		"total":       total,
		"total_pages": totalPages,
		"rows":        invsRes,
		"keyword":     keyword,
	}))
}
