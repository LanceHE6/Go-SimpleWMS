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

	err := db.Select("*").Find(&gts).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the list of goods type",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	var gtsRes []model.GoodsType
	for _, gt := range gts {
		gtsRes = append(gtsRes, gt)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Get goods type list successfully",
		"rows":    gtsRes,
		"code":    201,
	})
}
