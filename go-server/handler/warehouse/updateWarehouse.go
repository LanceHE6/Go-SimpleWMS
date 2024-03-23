package warehouse

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateWarehouse(context *gin.Context) {
	wid := context.PostForm("wid")
	warehouseName := context.PostForm("name")
	comment := context.PostForm("comment")

	if wid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "wid is required"})
		return
	}
	if warehouseName == "" && comment == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "name or comment is required"})
		return
	}

	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 判断该仓库是否已存在
	var registered int
	err := tx.QueryRow("SELECT count(name) FROM warehouse WHERE wid=?", wid).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the number of warehouses for this wid"})
		return
	}
	if registered == 0 {
		context.JSON(http.StatusForbidden, gin.H{"message": "The warehouse does not exist"})
		return
	}

	// 更新仓库

	if warehouseName == "" {
		_, err = tx.Exec("UPDATE warehouse SET comment=? WHERE wid=?", comment, wid)
	} else {
		// 判断该仓库名是否已存在
		var registered int
		err = tx.QueryRow("SELECT count(name) FROM warehouse WHERE name=?", warehouseName).Scan(&registered)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the number of warehouses for this warehouse_name"})
			return
		}
		if registered >= 1 {
			context.JSON(http.StatusForbidden, gin.H{"message": "The warehouse name already exists"})
			return
		}

		if comment == "" {
			_, err = tx.Exec("UPDATE warehouse SET name=? WHERE wid=?", warehouseName, wid)
		} else {
			_, err = tx.Exec("UPDATE warehouse SET name=?, comment=? WHERE wid=?", warehouseName, comment, wid)
		}
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update the warehouse"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit the transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Warehouse updated successfully"})
}
