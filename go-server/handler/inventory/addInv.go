package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type addInvRequest struct {
	Date         string `json:"date" form:"date"`                                        // 单据日期
	Number       string `json:"number" form:"number"`                                    // 单号
	Department   string `json:"department" form:"department"`                            // 单据所属部门
	GoodsList    string `json:"goods_list" form:"goods_list" binding:"required"`         // 单据包含的货品
	InventoryTpe string `json:"inventory_type" form:"inventory_type" binding:"required"` // 出入库类型
	Operator     string `json:"operator" form:"operator" binding:"required"`             // 经办人
	Comment      string `json:"comment" form:"comment"`                                  // 备注
	Manufacturer string `json:"manufacturer" form:"manufacturer"`                        // 制造商
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
	Date := data.Date
	Number := data.Number
	Department := data.Department
	var GoodsList model.GoodsList
	err := json.Unmarshal([]byte(data.GoodsList), &GoodsList)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The format of the goods_list is incorrect",
			"code":    402,
			"detail":  err.Error(),
		})
		return
	}
	InventoryType := data.InventoryTpe
	Sid := data.Operator
	Comment := data.Comment
	Manufacturer := data.Manufacturer

	db := myDb.GetMyDbConnection()

	// 出入库类型存在性判断
	var iType model.InventoryType
	err = db.Model(&model.InventoryType{}).Where("itid=?", InventoryType).First(&iType).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The inventory type does not exist",
			"code":    402,
		})
		return
	}

	Iid := "i" + utils.GenerateUuid(8) // 转换为 8 位字符串

	// 单号为空时构建单号
	if Number == "" {
		nowTime := time.Now().Format("200601021504")
		if iType.Type == 1 {
			Number = "I" + nowTime
		} else {
			Number = "O" + nowTime
		}
	}

	var parsedDate time.Time
	if Date == "" {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		parsedDate, _ = time.ParseInLocation("2006-01-02 15:04:05", nowTime, time.Local)

	} else {
		parsedDate, err = time.ParseInLocation("2006-01-02 15:04:05", Date, time.Local)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "The format of the date is incorrect",
				"code":    403,
				"detail":  err.Error(),
			})
			return
		}
	}

	var inventory = model.Inventory{
		Iid:           Iid,
		Date:          parsedDate,
		Department:    Department,
		GoodsList:     GoodsList,
		Number:        Number,
		InventoryType: InventoryType,
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
	for _, goodsOrder := range GoodsList {
		var goods model.Goods
		db.Model(model.Goods{}).Where("gid = ?", goodsOrder.Goods).First(&goods)
		if iType.Type == 1 {
			goods.Quantity += goodsOrder.Amount
		} else {
			goods.Quantity -= goodsOrder.Amount
		}

		err = db.Model(model.Goods{}).Where("gid=?", goodsOrder.Goods).Updates(&goods).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot update goods",
				"code":   502,
				"detail": err.Error(),
			})
			return
		}

	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Inventory added successfully",
		"code":    201,
	})
}
