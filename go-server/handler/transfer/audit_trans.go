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
)

type auditTransRequest struct {
	Tid          string `json:"tid" form:"tid" binding:"required"`
	Passed       *bool  `json:"passed" form:"passed" binding:"required"`
	AuditComment string `json:"audit_comment" form:"audit_comment"`
}

func AuditTrans(context *gin.Context) {
	var data auditTransRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}

	tid := data.Tid
	myClaims, _ := utils.GetUserInfoByContext(context)
	uid := myClaims.Uid
	passed := *data.Passed
	auditComment := data.AuditComment

	db := my_db.GetMyDbConnection()
	var trans model.Transfer
	notExist := db.Model(model.Transfer{}).Where("tid = ?", tid).First(&trans).RecordNotFound()
	if notExist {
		context.JSON(http.StatusBadRequest, response.Response(402, "Transfer not exist", nil))
		return
	}
	// 判断是否已经审核
	if trans.Audited {
		context.JSON(http.StatusOK, response.Response(202, "Transfer already audited", nil))
		return
	}
	nowDateTime := utils.GetNowDateTime()

	var updateData = map[string]interface{}{
		"audited":       true,
		"passed":        passed,
		"audited_time":  nowDateTime,
		"auditor":       uid,
		"audit_comment": auditComment,
	}

	if passed {
		goodsList, _ := json.Marshal(trans.GoodsList)
		var outAddInvRequest = inventory.AddInvRequest{
			Date:         nowDateTime,
			Number:       trans.Number + "CK",
			Department:   "",
			GoodsList:    string(goodsList),
			Warehouse:    trans.SourWarehouse,
			InventoryTpe: "_default-db2_",
			Comment:      trans.Comment,
			Operator:     trans.Operator,
		}
		outIid, code, resp := inventory.DoAddInv(context, outAddInvRequest, true)
		if outIid == "" {
			context.JSON(code, resp)
			return
		} else if outIid == "UPDATE_ERROR" {
			return
		}

		var inAddInvRequest = inventory.AddInvRequest{
			Date:         nowDateTime,
			Number:       trans.Number + "RK",
			Department:   "",
			GoodsList:    string(goodsList),
			Warehouse:    trans.DestWarehouse,
			InventoryTpe: "_default-db1_",
			Comment:      trans.Comment,
			Operator:     trans.Operator,
		}
		inIid, code, resp := inventory.DoAddInv(context, inAddInvRequest, true)
		if inIid == "" {
			context.JSON(code, resp)
			return
		} else if inIid == "UPDATE_ERROR" {
			return
		}

		updateData["out_iid"] = outIid
		updateData["in_iid"] = inIid
	}

	db.Model(model.Transfer{}).Where("tid = ?", tid).Updates(updateData)
	context.JSON(http.StatusOK, response.Response(200, "Transfer audited", nil))
	return

}
