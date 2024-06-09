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

	var goodsRes []gin.H
	for _, g := range goods {
		goodsMeta := gin.H{
			"created_at":   g.CreatedAt,
			"update_at":    g.UpdatedAt,
			"gid":          g.Gid,
			"goods_code":   g.GoodsCode,
			"name":         g.Name,
			"model":        g.Model,
			"goods_type":   g.GoodsType,
			"manufacturer": g.Manufacturer,
			"unit":         g.Unit,
			"unit_price":   g.UnitPrice,
			"image":        g.Image,
		}
		goodsRes = append(goodsRes, goodsMeta)
	}

	context.JSON(http.StatusOK, gin.H{
		"code":        201,
		"message":     "Query successfully",
		"page":        page,
		"page_size":   limit,
		"total":       total,
		"total_pages": totalPages,
		"keyword":     keyword,
		"rows":        goodsRes,
	})

}
