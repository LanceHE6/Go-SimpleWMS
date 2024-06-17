package stock

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetStockRequest(context *gin.Context) {
	// 获取库存信息
	goods := context.Query("goods")
	warehouse := context.Query("warehouse")
	db := myDb.GetMyDbConnection()
	if warehouse == "" && goods == "" {
		// 获取所有的仓库的所有物品库存信息
		type goodsQuery struct {
			Gid      string
			Quantity float64
		}
		type goodsResponse struct {
			Goods    model.Goods `json:"goods"`
			Quantity float64     `json:"quantity"`
		}
		var result []goodsQuery
		db.Table("stocks").
			Select("goods AS gid, SUM(quantity) AS quantity").
			Group("goods").
			Find(&result)
		var goodsResponseList []goodsResponse
		for _, v := range result {
			var goods model.Goods
			db.Model(model.Goods{}).Where("gid = ?", v.Gid).First(&goods)
			goodsResponseList = append(goodsResponseList, goodsResponse{
				Goods:    goods,
				Quantity: v.Quantity,
			},
			)
		}

		context.JSON(http.StatusOK, response.Response(201, "Get stock success", gin.H{
			"rows":  goodsResponseList,
			"total": len(goodsResponseList),
		}))
		return
	}

	type warehouseQuery struct {
		CreatedAT time.Time   `json:"created_at"`
		UpdatedAT time.Time   `json:"updated_at"`
		Goods     model.Goods `json:"goods"`
		Warehouse string      `json:"warehouse"`
		Quantity  float64     `json:"quantity"`
	}
	// 获取仓库的库存信息
	if goods == "" {

		var stocks []model.Stock
		err := db.Model(model.Stock{}).Where("warehouse = ?", warehouse).Find(&stocks).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Get stock failed", err.Error()))
			return
		}

		var warehouseResponseList []warehouseQuery
		for _, stock := range stocks {
			var goods model.Goods
			db.Model(model.Goods{}).Where("gid = ?", stock.Goods).First(&goods)
			warehouseResponseList = append(warehouseResponseList, warehouseQuery{
				CreatedAT: stock.CreatedAt,
				UpdatedAT: stock.UpdatedAt,
				Goods:     goods,
				Warehouse: stock.Warehouse,
				Quantity:  stock.Quantity,
			},
			)
		}

		context.JSON(http.StatusOK, response.Response(202, "Get stock success", gin.H{
			"rows":  warehouseResponseList,
			"total": len(warehouseResponseList),
		}))
		return
	}
	// 获取商品的总库存信息
	if warehouse == "" {
		var stocks []model.Stock
		err := db.Model(model.Stock{}).Where("goods = ?", goods).Find(&stocks).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(502, "Get stock failed", err.Error()))
			return
		}
		var goodsRes model.Goods
		db.Model(model.Goods{}).Where("gid = ?", goods).First(&goodsRes)
		// 计算总库存
		var totalQuantity float64
		var resp []warehouseQuery
		for _, stock := range stocks {
			resp = append(resp, warehouseQuery{
				CreatedAT: stock.CreatedAt,
				UpdatedAT: stock.UpdatedAt,
				Goods:     goodsRes,
				Warehouse: stock.Warehouse,
				Quantity:  stock.Quantity,
			},
			)
			totalQuantity += stock.Quantity
		}

		context.JSON(http.StatusOK, response.Response(203, "Get stock success", gin.H{
			"total_quantity": totalQuantity,
			"rows":           resp,
		}))
		return
	}
	// 获取指定仓库的指定商品的库存信息
	quantity := GetStock(warehouse, goods)
	var g model.Goods
	db.Model(model.Goods{}).Where("gid = ?", goods).First(&g)
	context.JSON(http.StatusOK, response.Response(204, "Get stock success", gin.H{
		"goods":     g,
		"warehouse": warehouse,
		"quantity":  quantity,
	}))
}
