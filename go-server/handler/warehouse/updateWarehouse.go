package warehouse

import (
	"Go_simpleWMS/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateWarehouseRequest struct {
	Wid     string `json:"wid" form:"wid" binding:"required"`
	Name    string `json:"name" form:"name"`
	Comment string `json:"comment" form:"comment"`
	Manager string `json:"manager" form:"manager"`
	Status  string `json:"status" form:"status"`
}

func UpdateWarehouse(context *gin.Context) {
	var data updateWarehouseRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "wid is required"})
		return
	}
	wid := data.Wid
	warehouseName := data.Name
	comment := data.Comment
	manager := data.Manager
	status := data.Status

	if warehouseName == "" && comment == "" && manager == "" && status == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "One of name, comment, manager and status is required"})
		return
	}

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
		})
		return
	}

	// 判断该仓库是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM warehouse WHERE wid=?", wid).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of warehouses for this wid",
			"detail": err.Error(),
		})
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
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot get the number of warehouses for this warehouse_name",
				"detail": err.Error(),
			})
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
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update the warehouse",
			"detail": err.Error(),
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Warehouse updated successfully"})
}
