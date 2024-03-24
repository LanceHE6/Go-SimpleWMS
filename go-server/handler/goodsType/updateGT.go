package goodsType

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateGoodsType(context *gin.Context) {
	GTid := context.PostForm("gtid")
	GTName := context.PostForm("name")
	typeCode := context.PostForm("type_code")

	if GTid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "GTid is required"})
		return
	}
	if GTName == "" && typeCode == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "name or type_code is required"})
		return
	}

	tx, _ := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin transaction"})
		return
	}

	// 判断该类型是否已存在
	var registered int
	err := tx.QueryRow("SELECT count(name) FROM goods_type WHERE gtid=?", GTid).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the number of goods type for this gtid"})
		return
	}
	if registered == 0 {
		context.JSON(http.StatusForbidden, gin.H{"message": "The goods type does not exist"})
		return
	}

	// 更新仓库

	if GTName == "" {
		_, err = tx.Exec("UPDATE goods_type SET type_code=? WHERE gtid=?", typeCode, GTid)
	} else {
		// 判断该仓库名是否已存在
		var registered int
		err = tx.QueryRow("SELECT count(name) FROM goods_type WHERE name=?", GTName).Scan(&registered)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get the number of goods type for this type name"})
			return
		}
		if registered >= 1 {
			context.JSON(http.StatusForbidden, gin.H{"message": "The type name already exists"})
			return
		}

		if typeCode == "" {
			_, err = tx.Exec("update goods_type set name=? where gtid=?", GTName, GTid)
		} else {
			_, err = tx.Exec("update goods_type set name=?, type_code=? where gtid=?", GTName, typeCode, GTid)
		}
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update the goods type"})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot commit the transaction"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Goods type updated successfully"})
}
