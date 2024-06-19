package transfer

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/handler/inventory"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type revokeAuditRequest struct {
	Tid string `json:"tid" form:"tid" binding:"required"`
}

func RevokeAudit(context *gin.Context) {
	var data revokeAuditRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}

	tid := data.Tid

	db := myDb.GetMyDbConnection()
	var trans model.Transfer
	notExist := db.Model(model.Transfer{}).Where("tid = ?", tid).First(&trans).RecordNotFound()
	if notExist {
		context.JSON(http.StatusBadRequest, response.Response(402, "Transfer not exist", nil))
		return
	}
	if !trans.Audited {
		context.JSON(http.StatusOK, response.Response(202, "The transfer has not been audited", nil))
		return
	}

	var updateData = map[string]interface{}{
		"audited":       false,
		"passed":        nil,
		"audited_time":  nil,
		"auditor":       nil,
		"audit_comment": "",
	}

	if trans.Passed {
		// 分别删除两个出入库单
		inIid := trans.InIid
		outIid := trans.OutIid

		code, resp := inventory.DoDeleteInv(inventory.DelInvRequest{Iid: outIid})
		if code != http.StatusOK {
			context.JSON(code, resp)
			return
		}

		code, resp = inventory.DoDeleteInv(inventory.DelInvRequest{Iid: inIid})
		if code != http.StatusOK {
			context.JSON(code, resp)
			return
		}

	}

	db.Model(model.Transfer{}).Where("tid = ?", tid).Updates(updateData)
	context.JSON(http.StatusOK, response.Response(200, "Audited revocation successful", nil))
	return
}
