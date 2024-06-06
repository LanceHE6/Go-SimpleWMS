package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type updateInvRequest struct {
	Iid          string `json:"iid" form:"iid" binding:"required"`
	Date         string `json:"date" form:"date"`                     // 单据日期
	Number       string `json:"number" form:"number"`                 // 单号
	Department   string `json:"department" form:"department"`         // 单据所属部门
	GoodsList    string `json:"goods_list" form:"goods_list"`         // 单据包含的货品
	InventoryTpe string `json:"inventory_type" form:"inventory_type"` // 出入库类型
	Warehouse    string `json:"warehouse" form:"warehouse"`           // 所属仓库
	Operator     string `json:"operator" form:"operator"`             // 经办人
	Comment      string `json:"comment" form:"comment"`               // 备注
	Manufacturer string `json:"manufacturer" form:"manufacturer"`     // 制造商
}

func UpdateInv(context *gin.Context) {
	var data updateInvRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	// 获取请求参数
	iid := data.Iid
	date := data.Date
	number := data.Number
	department := data.Department
	var goodsList model.GoodsList
	err := json.Unmarshal([]byte(data.GoodsList), &goodsList)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "The format of the goods_list is incorrect",
			"code":    402,
			"detail":  err.Error(),
		})
		return
	}
	inventoryTpe := data.InventoryTpe
	warehouse := data.Warehouse
	operator := data.Operator
	comment := data.Comment
	manufacturer := data.Manufacturer

	parsedDate, _ := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
	// 构造更新数据
	var inv = map[string]interface{}{
		"date":           parsedDate,
		"number":         number,
		"department":     department,
		"goods_list":     goodsList,
		"inventory_type": inventoryTpe,
		"warehouse":      warehouse,
		"operator":       operator,
		"comment":        comment,
		"manufacturer":   manufacturer,
	}

	db := myDb.GetMyDbConnection()
	// 获取原来的出入库单
	var oldInv model.Inventory
	db.Where("iid = ?", iid).First(&oldInv)

	// 更新出入库单
	err = db.Model(&model.Inventory{}).Where("iid = ?", iid).Updates(inv).Error
	// 单号重复判断
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":  "The Number is already exists",
				"detail": err.Error(),
				"code":   404,
			})
			return
		}
	}

	// 根据原来的出入库单回退对应的货品数量
	for _, goodsInfo := range oldInv.GoodsList {
		var goods = model.Goods{
			Gid: goodsInfo.Goods,
		}
		db.Where("gid = ?", goodsInfo.Goods).First(&goods)
		// 获取出入库类型
		itid := oldInv.InventoryType
		var inventoryType model.InventoryType
		db.Where("itid = ?", itid).First(&inventoryType)
		// 根据出入库类型回退数量
		if inventoryType.Type == 1 {
			goods.Quantity -= goodsInfo.Amount
		} else {
			goods.Quantity += goodsInfo.Amount
		}
		db.Model(model.Goods{}).Where("gid=?", goods.Gid).Update(&goods)
	}

	// 获取新的出入单
	var newInv model.Inventory
	db.Where("iid = ?", iid).First(&newInv)

	// 更新货品数量
	for _, goodsInfo := range newInv.GoodsList {
		var goods = model.Goods{
			Gid: goodsInfo.Goods,
		}
		db.Where("gid = ?", goodsInfo.Goods).First(&goods)
		// 获取出入库类型
		itid := newInv.InventoryType
		var inventoryType model.InventoryType
		db.Where("itid = ?", itid).First(&inventoryType)
		// 根据出入库类型更新数量
		if inventoryType.Type == 1 {
			goods.Quantity += goodsInfo.Amount

		} else {
			goods.Quantity -= goodsInfo.Amount
		}

		db.Model(model.Goods{}).Where("gid=?", goods.Gid).Update(&goods)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Inventory updated successfully",
		"code":    201,
	})

}
