package group

import (
	"Go_simpleWMS/handler/upload"
	"github.com/gin-gonic/gin"
)

func UploadGroup(ginApi *gin.RouterGroup) {
	userGroup := ginApi.Group("/upload")

	userGroup.POST("/goods_img", func(c *gin.Context) {
		upload.GoodsAttachmentUpload(c, upload.IMAGE)
	})
	userGroup.POST("/goods_file", func(c *gin.Context) {
		upload.GoodsAttachmentUpload(c, upload.FILE)
	})
}
