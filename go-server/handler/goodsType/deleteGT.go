package goodsType

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteGoodsTypeRequest struct {
	GTid string `json:"gtid" form:"gtid" binding:"required"`
}

func DeleteGoodsType(context *gin.Context) {
	var data deleteGoodsTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	gtid := data.GTid

	db := myDb.GetMyDbConnection()

	// 删除仓库
	err := db.Delete(&model.GoodsType{}, "gtid=?", gtid).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the Goods type",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Goods type deleted successfully",
		"code":    201,
	})
}
