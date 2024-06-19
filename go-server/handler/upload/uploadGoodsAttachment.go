package upload

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"Go_simpleWMS/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
)

type ContentType int

const (
	IMAGE ContentType = 1
	FILE  ContentType = 2
)

var imageExt = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	".svg":  true,
}

var fileExt = map[string]bool{
	".doc":  true,
	".docx": true,
	".pdf":  true,
	".xls":  true,
	".xlsx": true,
	".ppt":  true,
	".pptx": true,
	".txt":  true,
	".zip":  true,
	".rar":  true,
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	".svg":  true,
}

type uploadGoodsImageRequest struct {
	Goods string `json:"goods" form:"goods" binding:"required"`
}

// GoodsAttachmentUpload 上传货品相关附件
func GoodsAttachmentUpload(context *gin.Context, contentType ContentType) {
	var data uploadGoodsImageRequest
	fileDir := ""
	if contentType == IMAGE {
		fileDir = "static/res/goodsImage"
	}
	if contentType == FILE {
		fileDir = "static/res/goodsFile"
	}
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, response.MissingParamsResponse(err))
		return
	}
	// 判断本地静态文件夹是否存在
	_, err := os.Stat(fileDir)
	if err != nil {
		// 创建文件夹
		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(501, "Failed to create directory", err.Error()))
			return
		}
	}
	requestKey := ""
	if contentType == IMAGE {
		requestKey = "image"
	}
	if contentType == FILE {
		requestKey = "file"
	}
	form, err := context.MultipartForm()
	if err != nil {
		context.JSON(http.StatusBadRequest, response.Response(402, "Failed to get multipart form", nil))
		return
	}
	fileHeaders := form.File[requestKey]
	gid := data.Goods

	// 获取数据库连接
	db := myDb.GetMyDbConnection()
	// 查询货品信息
	var goods model.Goods
	err = db.Model(&model.Goods{}).Where("gid=?", gid).First(&goods).Error
	if err != nil {
		context.JSON(http.StatusOK, response.Response(403, "Goods not found", nil))
		return
	}

	var res []map[string]interface{}
	var filePaths model.FileList

	// 处理每个文件
	var count = 0
	for _, header := range fileHeaders {
		count += 1
		file, err := header.Open()
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(502, "Failed to open file", err.Error()))
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("Failed to close file:", err.Error())
			}
		}(file)

		fileName := header.Filename
		// 获取文件后缀
		extString := path.Ext(fileName)
		var ext map[string]bool
		if contentType == IMAGE {
			ext = imageExt
		}
		if contentType == FILE {
			ext = fileExt
		}

		if !ext[extString] {
			context.JSON(http.StatusBadRequest, response.Response(404, "Invalid file type", nil))
			return
		}
		// 构造文件名
		fileName = "goods_" + gid + "_" + strconv.Itoa(count) + extString
		var filePath model.File
		filePath.Path = fileDir + "/" + fileName

		out, err := os.Create(filePath.Path)
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(503, "Failed to create file", err.Error()))
			return
		}
		defer func(out *os.File) {
			err := out.Close()
			if err != nil {
				fmt.Println("Failed to close file:", err.Error())
			}
		}(out)

		// 保存文件
		_, ioErr := io.Copy(out, file)
		if ioErr != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(504, "Failed to save file", ioErr.Error()))
			return
		}
		filePaths = append(filePaths, filePath)
		res = append(res, map[string]interface{}{
			"gid":  gid,
			"path": filePath.Path,
			"name": fileName,
		})
	}

	// 更新数据库
	if contentType == IMAGE {
		var goods model.Goods
		db.Model(model.Goods{}).Where("gid = ?", gid).First(&goods)
		goods.Images = filePaths
		db.Model(model.Goods{}).Where("gid = ?", gid).Updates(&goods)
	}
	if contentType == FILE {
		var goods model.Goods
		db.Model(model.Goods{}).Where("gid = ?", gid).First(&goods)
		goods.Files = filePaths
		db.Model(model.Goods{}).Where("gid = ?", gid).Updates(&goods)
	}

	context.JSON(http.StatusOK, response.Response(200, "Upload successful", gin.H{
		"data": res,
	}))
}
