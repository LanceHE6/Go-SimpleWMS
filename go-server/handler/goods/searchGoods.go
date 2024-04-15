package goods

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func SearchGoods(context *gin.Context) {
	var goods []model.Goods
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	name := context.Query("name")
	gModel := context.Query("model")
	goodsType := context.Query("goods_type")
	warehouse := context.Query("warehouse")
	manufacturer := context.Query("manufacturer")
	quantity, _ := strconv.Atoi(context.DefaultQuery("quantity", "0"))
	keyword := context.Query("keyword")

	query := myDb.GetMyDbConnection()
	// 计算偏移量
	offset := (page - 1) * limit

	if name != "" {
		query = query.Where("name = ?", name)
	}
	if gModel != "" {
		query = query.Where("model = ?", gModel)
	}
	if goodsType != "" {
		query = query.Where("goods_type = ?", goodsType)
	}
	if warehouse != "" {
		query = query.Where("warehouse = ?", warehouse)
	}
	if manufacturer != "" {
		query = query.Where("manufacturer = ?", manufacturer)
	}
	if quantity != 0 {
		query = query.Where("quantity = ?", quantity)
	}
	if keyword != "" {
		query = query.Where("name LIKE ? OR model LIKE ? OR goods_type LIKE ? OR warehouse LIKE ? OR manufacturer LIKE ? OR quantity LIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%",
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
		context.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
			"code":  401,
		})
		return
	}
	if len(goods) == 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":        202,
			"message":     "No data",
			"page":        page,
			"page_size":   limit,
			"total":       total,
			"total_pages": totalPages,
			"rows":        goods,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":        201,
		"message":     "Query successfully",
		"page":        page,
		"page_size":   limit,
		"total":       total,
		"total_pages": totalPages,
		"keyword":     keyword,
		"rows":        goods,
	})

}
