package op_log

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func SearchOPLog(context *gin.Context) {
	var opLogs []model.OPLog
	var total int64
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	uid := context.Query("uid")
	resource := context.Query("resource")
	status := context.Query("status")

	query := my_db.GetMyDbConnection().Table("op_logs")

	if uid != "" {
		query = query.Where("uid = ?", uid)
	}
	if resource != "" {
		query = query.Where("resource = ?", resource)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var totalPages = 0
	// 获取总记录数
	query.Count(&total)
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
	result := query.Find(&opLogs)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Database query error", result.Error.Error()))
		return
	}
	if len(opLogs) == 0 {
		context.JSON(http.StatusOK, response.Response(202, "No data", gin.H{
			"page":        page,
			"page_size":   limit,
			"total":       total,
			"total_pages": totalPages,
			"rows":        opLogs,
		}))
		return
	}

	context.JSON(http.StatusOK, response.Response(201, "Query successfully", gin.H{
		"page":        page,
		"page_size":   limit,
		"total":       total,
		"total_pages": totalPages,
		"rows":        opLogs,
	}))

}
