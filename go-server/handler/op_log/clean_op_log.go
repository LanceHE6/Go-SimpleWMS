package op_log

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CleanOPLog(context *gin.Context) {
	db := my_db.GetMyDbConnection()
	db.Delete(&model.OPLog{})
	context.JSON(http.StatusOK, response.Response(200, "Clear operation log successful", nil))
}
