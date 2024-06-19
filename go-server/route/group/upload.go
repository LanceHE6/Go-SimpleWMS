package group

import (
	"Go_simpleWMS/handler/upload"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
)

func UploadGroup(ginApi *gin.RouterGroup) {
	userGroup := ginApi.Group("/upload")

	userGroup.POST("/goods_img",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("货品", "上传货品图片"),
		func(c *gin.Context) {
			upload.GoodsAttachmentUpload(c, upload.IMAGE)
		})
	userGroup.POST("/goods_file",
		utils.IsAdminMiddleware(),
		utils.OPLoggerMiddleware("货品", "上传货品附件"),
		func(c *gin.Context) {
			upload.GoodsAttachmentUpload(c, upload.FILE)
		})
}
