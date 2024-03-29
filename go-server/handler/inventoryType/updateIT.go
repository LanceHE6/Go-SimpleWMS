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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "itid is required",
			"code":    401,
		})
		return
	}
	ITid := data.ITid
	ITName := data.Name
	typeCode := data.TypeCode

	if ITName == "" && typeCode == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Name or type_code is required",
			"code":    402,
		})
		return
	}

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
			"code":   501,
		})
		return
	}

	// 判断该类型是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM inventory_type WHERE itid=?", ITid).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of inventory type for this itid",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	if registered == 0 {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The inventory type does not exist",
			"code":    403,
		})
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
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot get the number of inventory type for this type name",
				"detail": err.Error(),
				"code":   503,
			})
			return
		}
		if registered >= 1 {
			context.JSON(http.StatusForbidden, gin.H{
				"message": "The type name already exists",
				"code":    404,
			})
			return
		}

		if typeCode == "" {
			_, err = tx.Exec("update inventory_type set name=? where itid=?", ITName, ITid)
		} else {
			_, err = tx.Exec("update inventory_type set name=?, type_code=? where itid=?", ITName, typeCode, ITid)
		}
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update the inventory type",
			"detail": err.Error(),
			"code":   504,
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
			"code":   505,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Inventory type updated successfully",
		"code":    201,
	})
}
