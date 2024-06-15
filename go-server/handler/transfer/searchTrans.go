package transfer

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

func SearchTrans(context *gin.Context) {
	var transs []model.Transfer
	var total int64
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))

	goods := context.Query("goods")
	number := context.Query("number")
	sourWarehouse := context.Query("source_warehouse")
	destWarehouse := context.Query("Destination_warehouse")
	amount, _ := strconv.Atoi(context.DefaultQuery("amount", "0"))
	operator := context.Query("operator")
	comment := context.Query("comment")
	date := context.Query("date")
	auditedQ := context.Query("audited")
	passedQ := context.Query("passed")
	keyword := context.Query("keyword")

	query := myDb.GetMyDbConnection().Table("transfers")
	// 依据参数构造查询
	if goods != "" {
		query = query.Where(query.Where("JSON_CONTAINS(transfers.goods_list, JSON_OBJECT('Goods', ?))", goods))
	}
	if number != "" {
		query = query.Where("number = ?", number)
	}
	if sourWarehouse != "" {
		query = query.Where("sources_warehouse = ?", sourWarehouse)
	}
	if destWarehouse != "" {
		query = query.Where("destination_warehouse = ?", destWarehouse)
	}
	if amount != 0 {
		query = query.Where("JSON_CONTAINS(transfers.goods_list, JSON_OBJECT('Amount', ?))", amount)
	}
	if operator != "" {
		query = query.Where("operator = ?", operator)
	}
	if comment != "" {
		query = query.Where("comment = ? OR JSON_CONTAINS(goods_list, JSON_OBJECT('Comment', ?))", comment, comment)
	}
	if date != "" {
		// 将 created_at 转换为日期格式，并过滤出当天的记录
		parsedDate, err := time.Parse("2006-01-02", date)
		fmt.Println(parsedDate)
		if err == nil {
			startOfDay := parsedDate.Format("2006-01-02 00:00:00")
			endOfDay := parsedDate.Add(24 * time.Hour).Format("2006-01-02 15:04:05")
			query = query.Where("date BETWEEN ? AND ?", startOfDay, endOfDay)
		}
	}

	if auditedQ == "true" || auditedQ == "True" {
		query = query.Where("audited = ?", true)
	} else if auditedQ == "false" || auditedQ == "False" {
		query = query.Where("audited = ?", false)
	}

	if passedQ == "true" || passedQ == "True" {
		query = query.Where("passed = ?", true)
	} else if passedQ == "false" || passedQ == "False" {
		query = query.Where("passed = ?", false)
	}

	if keyword != "" {
		query = query.Where("goods_list LIKE ? OR number LIKE ? OR source_warehouse LIKE ? OR destination_warehouse LIKE ? OR operator LIKE ? OR comment LIKE ? OR date LIKE ?",
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
	result := query.Find(&transs)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Internal Server Error", result.Error.Error()))
		return
	}

	// 查询结果数量为0
	if len(transs) == 0 {
		context.JSON(http.StatusOK, response.Response(202, "No data", gin.H{
			"page":        page,
			"page_size":   limit,
			"total":       total,
			"total_pages": totalPages,
			"rows":        transs,
		}))
		return
	}

	var transRes []gin.H
	db := myDb.GetMyDbConnection()
	for _, transfer := range transs {

		var goodsList []gin.H
		// 构造货品清单返回体
		for _, goodsInfo := range transfer.GoodsList {
			var goods model.Goods
			db.Model(model.Goods{}).Where("gid = ?", goodsInfo.Goods).First(&goods)
			goodsList = append(goodsList, gin.H{
				"goods":   goods,
				"amount":  goodsInfo.Amount,
				"comment": goodsInfo.Comment,
			})
		}
		// 构造单个出入库单返回体
		transMeta := gin.H{
			"created_at":            transfer.CreatedAt,
			"update_at":             transfer.UpdatedAt,
			"tid":                   transfer.Tid,
			"date":                  transfer.Date,
			"number":                transfer.Number,
			"goods_list":            goodsList,
			"source_warehouse":      transfer.SourWarehouse,
			"destination_warehouse": transfer.DestWarehouse,
			"operator":              transfer.Operator,
			"comment":               transfer.Comment,
			"audited":               transfer.Audited,
			"auditor":               transfer.Auditor,
			"passed":                transfer.Passed,
			"audited_time":          transfer.AuditedTime,
			"audit_comment":         transfer.AuditComment,
		}
		transRes = append(transRes, transMeta)
	}

	context.JSON(http.StatusOK, response.Response(201, "Query successfully", gin.H{
		"page":        page,
		"page_size":   limit,
		"total":       total,
		"total_pages": totalPages,
		"rows":        transRes,
		"keyword":     keyword,
	}))
}
