package inventory

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	Warehouse := data.Warehouse
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
		Warehouse:     Warehouse,
		Number:        Number,
		InventoryType: InventoryType,
		Operator:      Sid,
		Comment:       Comment,
		Manufacturer:  Manufacturer,
	}

	UpdateStocks(GoodsList, Warehouse, iType, context, db)
	// 插入单据
	err = db.Model(model.Inventory{}).Create(&inventory).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot insert new inventory",
			"code":   501,
			"detail": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Inventory added successfully",
		"code":    201,
	})
}

func UpdateStocks(GoodsList model.GoodsList, Warehouse string, inventoryType model.InventoryType, context *gin.Context, db *gorm.DB) int {
	// 更新库存表
	for _, goodsOrder := range GoodsList {
		var stock model.Stock
		// 判断库存表中是否存在这个映射
		notExist := db.Model(model.Stock{}).Where("warehouse=? AND goods=?", Warehouse, goodsOrder.Goods).First(&stock).RecordNotFound()
		var newStock = model.Stock{
			Warehouse: Warehouse,
			Goods:     goodsOrder.Goods,
			Quantity:  goodsOrder.Amount,
		}
		// 根据出入库方向执行不同的操作
		if inventoryType.Type == 1 {
			if notExist {
				// 不存在就插入记录新增映射
				db.Model(model.Stock{}).Create(&newStock)
			} else {
				// 存在就更新记录
				db.Model(model.Stock{}).Where("warehouse=? AND goods=?", Warehouse, goodsOrder.Goods).Update("quantity", stock.Quantity+goodsOrder.Amount)
			}
		} else {
			if notExist {
				// 无仓库货品映射，无法执行出库操作
				context.JSON(http.StatusBadRequest, gin.H{
					"message": "The goods is not included in the warehouse and cannot be outbound",
					"code":    405,
					"detail":  "",
				})
				return 1
			} else {
				// 存在就更新记录
				if stock.Quantity < goodsOrder.Amount {
					context.JSON(http.StatusBadRequest, gin.H{
						"message": "Including outbound goods with insufficient stock",
						"code":    406,
						"detail":  "",
					})
					return 1
				} else {
					var updateData = map[string]interface{}{
						"quantity": stock.Quantity - goodsOrder.Amount,
					}
					db.Model(model.Stock{}).Where("warehouse=? AND goods=?", Warehouse, goodsOrder.Goods).Updates(updateData)
				}
			}
		}
	}
	return 0
}
