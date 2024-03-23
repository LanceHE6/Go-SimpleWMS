package warehouse

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteWarehouse(context *gin.Context) {
	wid := context.PostForm("wid")
	if wid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "wid is required"})
		return
	}

	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 删除仓库
	_, err := tx.Exec("DELETE FROM warehouse WHERE wid=?", wid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete the warehouse"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit the transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Warehouse deleted successfully"})
}
