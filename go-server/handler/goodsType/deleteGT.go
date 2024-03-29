package goodsType

import (
	"Go_simpleWMS/utils"
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
			"message": "GTid is required",
			"code":    401,
		})
		return
	}
	gtid := data.GTid

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	// 删除仓库
	_, err = tx.Exec("DELETE FROM goods_type WHERE gtid=?", gtid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the Goods type",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
			"code":   503,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Goods type deleted successfully",
		"code":    201,
	})
}
