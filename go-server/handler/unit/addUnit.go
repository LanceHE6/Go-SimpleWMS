package unit

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type addUnitRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func AddUnit(context *gin.Context) {
	var data addUnitRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	unitName := data.Name

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
	err = tx.QueryRow("SELECT count(name) FROM unit WHERE name=?", unitName).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of unit for this unit name",
			"detail": err.Error(),
			"code":   502,
		})
		return
	}
	if registered >= 1 {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The unit already exists",
			"code":    401,
		})
		return
	}

	newUnid := "un" + utils.GenerateUuid(8) // 转换为 8 位字符串

	addTime := time.Now().Unix()
	//增加仓库
	_, err = tx.Exec("INSERT INTO unit(unid, name, add_time) VALUES(?, ?, ?)", newUnid, unitName, addTime)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert the unit",
			"detail": err.Error(),
			"code":   505,
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
			"code":   506,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Unit added successfully",
		"code":    201,
	})
}
