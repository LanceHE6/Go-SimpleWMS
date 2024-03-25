package goodsType

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteGoodsType(context *gin.Context) {
	gtid := context.PostForm("gtid")
	if gtid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "gtid is required"})
		return
	}

	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 删除仓库
	_, err := tx.Exec("DELETE FROM goods_type WHERE gtid=?", gtid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete the Goods type"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit the transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Goods type deleted successfully"})
}
