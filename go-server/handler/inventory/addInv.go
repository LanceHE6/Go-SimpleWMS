package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type addInvRequest struct {
	Gid          string `json:"gid" form:"gid"`
	Name         string `json:"name" form:"name"`
	Amount       int    `json:"amount" form:"amount" binding:"required"`
	InventoryTpe string `json:"inventory_type" form:"inventory_type" binding:"required"`
	Warehouse    string `json:"warehouse" form:"warehouse" binding:"required"`
	Operator     string `json:"operator" form:"operator" binding:"required"`
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
	Name := data.Name
	Amount := data.Amount
	InventoryType := data.InventoryTpe
	Wid := data.Warehouse
	Sid := data.Operator
	Comment := data.Comment
	Manufacturer := data.Manufacturer

	db := myDb.GetMyDbConnection()

	var iType model.InventoryType
	err := db.Model(&model.InventoryType{}).Where("itid=?", InventoryType).First(&iType).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The inventory type does not exist",
			"code":    402,
		})
		return
	}

	Iid := "i" + utils.GenerateUuid(8) // 转换为 8 位字符串

	var goods model.Goods
	err = db.Model(&model.Goods{}).Where("gid=?", Gid).First(&goods).Error
	if err != nil {
		// 该货品不存在 判断出入库类型是否为入库
		// 1 为入库 2 为出库
		if iType.Type != 1 {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "The goods does not exist",
				"code":    403,
			})
			return
		}
		// 生成新的货品
		newGid := "g" + utils.GenerateUuid(8) // 转换为 8 位字符串
		if Name == "" {
			Name = "新添加货品"
		}

		goods = model.Goods{
			Gid:          newGid,
			Name:         Name,
			GoodsType:    "_default_",
			Warehouse:    Wid,
			Manufacturer: Manufacturer,
			Unit:         "_default_",
			Quantity:     0,
		}
		err = db.Create(&goods).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot insert the goods",
				"detail": err.Error(),
				"code":   501,
			})
			return
		}
	}

	goodsCode := goods.GoodsCode
	if goodsCode == "" {
		goodsCode = "G" + utils.GenerateUuid(4)
	}

	// 构建单号
	nowTime := time.Now().Format("200601021504")
	Number := ""
	if iType.Type == 1 {
		Number = "I" + nowTime + goodsCode
	} else {
		Number = "O" + nowTime + goodsCode
	}

	var inventory = model.Inventory{
		Iid:           Iid,
		Goods:         goods.Gid,
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
	if iType.Type == 1 {
		goods.Quantity += Amount
	} else {
		goods.Quantity -= Amount
	}

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
