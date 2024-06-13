package goods

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func SearchGoods(context *gin.Context) {
	var goods []model.Goods
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	gid := context.Query("gid")
	name := context.Query("name")
	gModel := context.Query("model")
	goodsType := context.Query("goods_type")
	manufacturer := context.Query("manufacturer")
	unitPrice, _ := strconv.ParseFloat(context.DefaultQuery("unit_price", "0"), 64)
	keyword := context.Query("keyword")

	query := myDb.GetMyDbConnection().Table("goods") //.Joins("left join warehouses on goods.warehouse = warehouses.wid").Where("warehouses.status = 1")

	// 计算偏移量
	offset := (page - 1) * limit

	if gid != "" {
		query = query.Where("goods.gid = ?", gid)
	}
	if name != "" {
		query = query.Where("goods.name = ?", name)
	}
	if gModel != "" {
		query = query.Where("goods.model = ?", gModel)
	}
	if goodsType != "" {
		query = query.Where("goods.goods_type = ?", goodsType)
	}
	if manufacturer != "" {
		query = query.Where("goods.manufacturer = ?", manufacturer)
	}
	if unitPrice != 0 {
		query = query.Where("goods.unit_price = ?", unitPrice)
	}
	if keyword != "" {
		query = query.Where("goods.name LIKE ? OR goods.model LIKE ? OR goods.goods_type LIKE ? OR goods.manufacturer LIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%")
	}

	// 获取总记录数
	var total int64
	query.Model(&model.Goods{}).Count(&total)

	// 计算总页数
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	// 设置分页参数
	query = query.Offset(offset).Limit(limit)

	result := query.Offset(offset).Limit(limit).Find(&goods)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Database query error", result.Error.Error()))
		return
	}
	if len(goods) == 0 {
		context.JSON(http.StatusOK, response.Response(202, "No data", gin.H{
			"page":        page,
			"page_size":   limit,
			"total":       total,
			"total_pages": totalPages,
			"rows":        goods,
		}))
		return
	}

	var goodsRes []model.Goods
	for _, g := range goods {
		goodsRes = append(goodsRes, g)
	}

	context.JSON(http.StatusOK, response.Response(201, "Query successfully", gin.H{
		"page":        page,
		"page_size":   limit,
		"total":       total,
		"total_pages": totalPages,
		"rows":        goods,
		"keyword":     keyword,
	}))

}
