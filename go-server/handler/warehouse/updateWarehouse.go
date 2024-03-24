package warehouse

import (
	"Go_simpleWMS/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateWarehouse(context *gin.Context) {
	wid := context.PostForm("wid")
	warehouseName := context.PostForm("name")
	comment := context.PostForm("comment")
	manager := context.PostForm("manager")
	status := context.PostForm("status")

	if wid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "wid is required"})
		return
	}
	if warehouseName == "" && comment == "" && manager == "" && status == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "One of name, comment, manager and status is required"})
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

		var updateSql = "Update warehouse Set "
		if comment != "" {
			updateSql += "comment= '" + comment + "',"
		}
		if manager != "" {
			updateSql += "manager= '" + manager + "',"
		}
		if status != "" {
			updateSql += "status= " + status + ","
		}
		updateSql = updateSql[:len(updateSql)-1] // 去掉最后一个逗号
		updateSql += " Where wid= '" + wid + "'"
		fmt.Println(updateSql)
		_, err = tx.Exec(updateSql)
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
