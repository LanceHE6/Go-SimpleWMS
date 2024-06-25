package transfer

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/handler/inventory"
	"Go_simpleWMS/utils"
	"Go_simpleWMS/utils/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type addTransRequest struct {
	Date          string `json:"date" form:"date"`                                                      // 单据日期
	Number        string `json:"number" form:"number"`                                                  // 单号
	GoodsList     string `json:"goods_list" form:"goods_list" binding:"required"`                       // 单据包含的货品
	SourWarehouse string `json:"source_warehouse" form:"source_warehouse" binding:"required"`           // 源仓库
	DestWarehouse string `json:"destination_warehouse" form:"destination_warehouse" binding:"required"` // 目标仓库
	Operator      string `json:"operator" form:"operator" binding:"required"`                           // 经办人
	Comment       string `json:"comment" form:"comment"`                                                // 备注
}

func AddTrans(context *gin.Context) {
	var data addTransRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	Date := data.Date
	Number := data.Number
	DestWarehouse := data.DestWarehouse
	SourWarehouse := data.SourWarehouse
	var GoodsList model.GoodsList
	err := json.Unmarshal([]byte(data.GoodsList), &GoodsList)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.Response(402, "The goods list is not in the correct format", nil))
		return
	}
	Sid := data.Operator
	Comment := data.Comment

	db := my_db.GetMyDbConnection()

	Tid := "t" + utils.GenerateUuid(8) // 转换为 8 位字符串

	// 单号为空时构建单号
	if Number == "" {
		nowTime := time.Now().Format("200601021504")
		Number = "DB" + nowTime + utils.GenerateUuid(4)
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

	var outAddInvRequest = inventory.AddInvRequest{
		Date:         Date,
		Number:       Number + "CK",
		Department:   "",
		GoodsList:    data.GoodsList,
		Warehouse:    SourWarehouse,
		InventoryTpe: "_default-db2_",
		Comment:      Comment,
		Operator:     Sid,
	}
	outIid, code, resp := inventory.DoAddInv(context, outAddInvRequest, false)
	if outIid == "" {
		context.JSON(code, resp)
		return
	} else if outIid == "UPDATE_ERROR" {
		return
	}

	var inAddInvRequest = inventory.AddInvRequest{
		Date:         Date,
		Number:       Number + "RK",
		Department:   "",
		GoodsList:    data.GoodsList,
		Warehouse:    DestWarehouse,
		InventoryTpe: "_default-db1_",
		Comment:      Comment,
		Operator:     Sid,
	}
	inIid, code, resp := inventory.DoAddInv(context, inAddInvRequest, false)
	if inIid == "" {
		context.JSON(code, resp)
		return
	} else if inIid == "UPDATE_ERROR" {
		return
	}

	var transfer = model.Transfer{
		Tid:           Tid,
		Date:          parsedDate,
		GoodsList:     GoodsList,
		SourWarehouse: SourWarehouse,
		DestWarehouse: DestWarehouse,
		Number:        Number,
		Operator:      Sid,
		Comment:       Comment,
	}
	// 插入单据
	err = db.Model(model.Transfer{}).Create(&transfer).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to add transfer", err.Error()))
		return
	}
	context.JSON(http.StatusOK, response.Response(200, "Add transfer successfully", nil))
}
