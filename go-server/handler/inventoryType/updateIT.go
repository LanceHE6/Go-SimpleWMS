package inventoryType

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateInventoryTypeRequest struct {
	ITid     string `json:"itid" form:"itid" binding:"required"`
	Name     string `json:"name" form:"name"`
	TypeCode string `json:"type_code" form:"type_code"`
}

func UpdateInventoryType(context *gin.Context) {
	var data updateInventoryTypeRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "itid is required"})
		return
	}
	ITid := data.ITid
	ITName := data.Name
	typeCode := data.TypeCode

	if ITName == "" && typeCode == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Name or type_code is required"})
		return
	}

	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 判断该类型是否已存在
	var registered int
	err := tx.QueryRow("SELECT count(name) FROM inventory_type WHERE itid=?", ITid).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the number of inventory type for this itid"})
		return
	}
	if registered == 0 {
		context.JSON(http.StatusForbidden, gin.H{"message": "The inventory type does not exist"})
		return
	}

	// 更新仓库

	if ITName == "" {
		_, err = tx.Exec("UPDATE inventory_type SET type_code=? WHERE itid=?", typeCode, ITid)
	} else {
		// 判断该类型名是否已存在
		var registered int
		err = tx.QueryRow("SELECT count(name) FROM inventory_type WHERE name=?", ITName).Scan(&registered)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the number of inventory type for this type name"})
			return
		}
		if registered >= 1 {
			context.JSON(http.StatusForbidden, gin.H{"message": "The type name already exists"})
			return
		}

		if typeCode == "" {
			_, err = tx.Exec("update inventory_type set name=? where itid=?", ITName, ITid)
		} else {
			_, err = tx.Exec("update inventory_type set name=?, type_code=? where itid=?", ITName, typeCode, ITid)
		}
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update the inventory type"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit the transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Inventory type updated successfully"})
}
