package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/handler/stock"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
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
	Warehouse    string `json:"warehouse" form:"warehouse" binding:"required"`           // 仓库
	InventoryTpe string `json:"inventory_type" form:"inventory_type" binding:"required"` // 出入库类型
	Operator     string `json:"operator" form:"operator" binding:"required"`             // 经办人
	Comment      string `json:"comment" form:"comment"`                                  // 备注
	Manufacturer string `json:"manufacturer" form:"manufacturer"`                        // 制造商
}

func AddInv(context *gin.Context) {
	var data addInvRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	Date := data.Date
	Number := data.Number
	Department := data.Department
	Warehouse := data.Warehouse
	var GoodsList model.GoodsList
	err := json.Unmarshal([]byte(data.GoodsList), &GoodsList)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.Response(402, "The goods list is not in the correct format", nil))
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
		context.JSON(http.StatusOK, response.Response(403, "The inventory type does not exist", nil))
		return
	}

	Iid := "i" + utils.GenerateUuid(8) // 转换为 8 位字符串

	// 单号为空时构建单号
	if Number == "" {
		nowTime := time.Now().Format("200601021504")
		if iType.Type == 1 {
			Number = "RK" + nowTime + utils.GenerateUuid(4)
		} else {
			Number = "CK" + nowTime + utils.GenerateUuid(4)
		}
	}

	var parsedDate time.Time
	if Date == "" {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		parsedDate, _ = time.ParseInLocation("2006-01-02 15:04:05", nowTime, time.Local)

	} else {
		parsedDate, err = time.ParseInLocation("2006-01-02 15:04:05", Date, time.Local)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.Response(404, "The date format is incorrect", nil))
			return
		}
	}
	// 构造更新前后库存数据
	var oldGoodsList model.GoodsList
	var newGoodsList model.GoodsList
	for _, g := range GoodsList {
		var oldGoodsOrder model.GoodsOrder
		var newGoodsOrder model.GoodsOrder
		oldStock := stock.GetStock(Warehouse, g.Goods)
		oldGoodsOrder.Goods = g.Goods
		newGoodsOrder.Goods = g.Goods
		oldGoodsOrder.Amount = oldStock
		if iType.Type == 1 {
			newGoodsOrder.Amount = oldStock + g.Amount
		} else {
			newGoodsOrder.Amount = oldStock - g.Amount
		}
		oldGoodsList = append(oldGoodsList, oldGoodsOrder)
		newGoodsList = append(newGoodsList, newGoodsOrder)
	}
	// 更新库存
	UpdateStocks(GoodsList, Warehouse, iType, context, db)

	var inventory = model.Inventory{
		Iid:           Iid,
		Date:          parsedDate,
		Department:    Department,
		GoodsList:     GoodsList,
		OldGoodsList:  oldGoodsList,
		NewGoodsList:  newGoodsList,
		Warehouse:     Warehouse,
		Number:        Number,
		InventoryType: InventoryType,
		Operator:      Sid,
		Comment:       Comment,
		Manufacturer:  Manufacturer,
	}
	// 插入单据
	err = db.Model(model.Inventory{}).Create(&inventory).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to add inventory", err.Error()))
		return
	}
	context.JSON(http.StatusOK, response.Response(200, "Add inventory successfully", nil))
}
