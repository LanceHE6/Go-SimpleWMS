package goods

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type deleteGoodsImgRequest struct {
	Goods     string `json:"goods" form:"goods" binding:"required"`
	ImageName string `json:"image_name" form:"image_name" binding:"required"`
}
type deleteGoodsFileRequest struct {
	Goods    string `json:"goods" form:"goods" binding:"required"`
	FileName string `json:"file_name" form:"file_name" binding:"required"`
}

type AttachmentType int

const (
	IMAGE AttachmentType = 1
	FILE  AttachmentType = 2
)

func DeleteGoodsAttachment(context *gin.Context, attachmentType AttachmentType) {

	if attachmentType == IMAGE {
		var data deleteGoodsImgRequest
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
			return
		}
		gid := data.Goods
		imgName := data.ImageName
		db := my_db.GetMyDbConnection()
		var goods model.Goods
		notExist := db.Model(model.Goods{}).Where("gid = ?", gid).First(&goods).RecordNotFound()
		if notExist {
			context.JSON(http.StatusBadRequest, response.Response(402, "The goods does not exist", nil))
			return
		}
		imageList := goods.Images
		// 构造新的imageList
		var newImageList []model.File
		for _, image := range imageList {
			if image.Name == imgName {
				// 删除本地的图片
				err := os.Remove(image.Path)
				if err != nil {
					context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Delete image failed", err.Error()))
					return
				}
			} else {
				newImageList = append(newImageList, image)
			}
		}
		// 更新数据库
		goods.Images = newImageList
		err := db.Save(&goods).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(502, "Delete image failed", err.Error()))
			return
		}

		context.JSON(http.StatusOK, response.Response(200, "Delete image successfully", nil))

	} else {
		var data deleteGoodsFileRequest
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
			return
		}
		gid := data.Goods
		fileName := data.FileName
		db := my_db.GetMyDbConnection()
		var goods model.Goods
		notExist := db.Model(model.Goods{}).Where("gid = ?", gid).First(&goods).RecordNotFound()
		if notExist {
			context.JSON(http.StatusBadRequest, response.Response(402, "The goods does not exist", nil))
			return
		}
		fileList := goods.Files
		// 构造新的fileList
		var newFileList []model.File
		for _, file := range fileList {
			if file.Name == fileName {
				// 删除本地的图片
				err := os.Remove(file.Path)
				if err != nil {
					context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Delete file failed", err.Error()))
					return
				}
			} else {
				newFileList = append(newFileList, file)
			}
		}
		// 更新数据库
		goods.Files = newFileList
		err := db.Save(&goods).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(502, "Delete file failed", err.Error()))
			return
		}

		context.JSON(http.StatusOK, response.Response(200, "Delete file successfully", nil))
	}

}
