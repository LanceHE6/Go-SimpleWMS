package goodsType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListGoodsType(context *gin.Context) {
	db := myDb.GetMyDbConnection()
	var gts []model.GoodsType

	err := db.Select([]string{"gtid", "name", "type_code", "created_at"}).Find(&gts).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of goods type",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	var gtsRes []gin.H
	for _, gt := range gts {

		gtRes := gin.H{
			"gtid":       gt.Gtid,
			"name":       gt.Name,
			"type_code":  gt.TypeCode,
			"created_at": gt.CreatedAt,
		}
		gtsRes = append(gtsRes, gtRes)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get goods type list successfully",
		"rows":    gtsRes,
		"code":    201,
	})
}
