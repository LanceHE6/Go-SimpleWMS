package goods

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteGoodsRequest struct {
	Gid string `json:"gid" form:"gid" binding:"required"`
}

func DeleteGoods(context *gin.Context) {
	var data deleteGoodsRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	gid := data.Gid

	db := myDb.GetMyDbConnection()

	// 删除
	var goods model.Goods
	db.Model(&model.Goods{}).Where("gid=?", gid).First(&goods)
	err := db.Delete(&goods).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot delete the Goods",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Goods deleted successfully",
		"code":    201,
	})
}
