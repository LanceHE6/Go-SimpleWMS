package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type addInvRequest struct {
	Gid          string `json:"gid" form:"gid" binding:"required"`
	Amount       int    `json:"amount" form:"amount" binding:"required"`
	InventoryTpe string `json:"inventory_type" form:"inventory_type" binding:"required"`
	Wid          string `json:"wid" form:"wid" binding:"required"`
	Sid          string `json:"sid" form:"sid" binding:"required"`
	Comment      string `json:"comment" form:"comment"`
	Manufacturer string `json:"manufacturer" form:"manufacturer"`
}

func AddInv(context *gin.Context) {
	var data addInvRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	Gid := data.Gid
	Amount := data.Amount
	InventoryType := data.InventoryTpe
	Wid := data.Wid
	Sid := data.Sid
	Comment := data.Comment
	Manufacturer := data.Manufacturer

	Iid := "i" + utils.GenerateUuid(8) // 转换为 8 位字符串

	db := myDb.GetMyDbConnection()
	var goods model.Goods
	err := db.Model(&model.Goods{}).Where("gid=?", Gid).First(&goods).Error
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "The goods does not exist",
			"code":    402,
		})
		return
	}
	goodsCode := goods.GoodsCode
	if goodsCode == "" {
		goodsCode = "G" + utils.GenerateUuid(4)
	}

	// 构建单号
	nowTime, _ := time.Parse("20060102150405", strconv.FormatInt(time.Now().Unix(), 10))
	Number := nowTime.String() + goodsCode

	var inventory = model.Inventory{
		Iid:           Iid,
		Goods:         Gid,
		Amount:        Amount,
		Number:        Number,
		InventoryType: InventoryType,
		Warehouse:     Wid,
		Operator:      Sid,
		Comment:       Comment,
		Manufacturer:  Manufacturer,
	}

	err = db.Model(model.Inventory{}).Create(&inventory).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert new inventory",
			"code":   501,
			"detail": err.Error(),
		})
		return
	}
	// 更新货品表
	goods.Quantity += Amount
	err = db.Model(model.Goods{}).Where("gid=?", goods.Gid).Updates(&goods).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update goods",
			"code":   502,
			"detail": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Inventory added successfully",
		"code":    201,
	})
}
