package transfer

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/handler/inventory"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteTransRequest struct {
	Tid string `json:"tid" form:"tid" binding:"required"`
}

func DeleteTrans(context *gin.Context) {
	var data deleteTransRequest
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

	// 审核通过执行了出入库操作的需回退
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

	// 删除Transfer
	db.Delete(&trans)

	context.JSON(http.StatusOK, response.Response(200, "Delete transfer successfully", nil))
}
