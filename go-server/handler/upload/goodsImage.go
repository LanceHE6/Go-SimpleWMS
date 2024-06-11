package upload

import (
	"Go_simpleWMS/database/model"
	"Go_simpleWMS/database/myDb"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

type uploadGoodsImageRequest struct {
	Gid string `json:"gid" form:"gid" binding:"required"`
}

func GoodsImageUpload(context *gin.Context) {
	var data uploadGoodsImageRequest
	fileDir := "static/res/goodsImage"
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters or incorrect format",
			"code":    401,
			"detail":  err.Error(),
		})
		return
	}
	// 判断本地静态文件夹是否存在
	_, err := os.Stat(fileDir)
	if err != nil {
		// 创建文件夹
		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Upload failed",
				"detail": err.Error(),
				"code":   501,
			})
			return
		}

	}
	file, header, err := context.Request.FormFile("image")
	gid := data.Gid

	if err == nil {
		fileName := header.Filename
		// 获取文件后缀
		extString := path.Ext(fileName)
		ext := map[string]bool{
			".png": true,
			".jpg": true,
		}
		if !ext[extString] {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Unsupported image format",
				"code":    402,
			})
			return
		}

		// 以时间戳命名
		timeStamp := time.Now().Unix()
		fileName = "goods_" + strconv.FormatInt(timeStamp, 10) + extString
		imagePath := fileDir + "/" + fileName
		// 更新数据库
		db := myDb.GetMyDbConnection()
		// gid 存在判断
		var goods model.Goods
		err := db.Model(&model.Goods{}).Where("gid=?", gid).Find(&goods).Error
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":  "The goods does not exist",
				"detail": err.Error(),
				"code":   403,
			})
			return
		}
		err = db.Model(&model.Goods{}).Where("gid=?", gid).Update("image", imagePath).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Updating database failed",
				"detail": err.Error(),
				"code":   504,
			})
			return
		}

		out, err := os.Create(imagePath)
		if err == nil {
			defer func(out *os.File) {
				err := out.Close()
				if err != nil {

				}
			}(out)

			// 保存文件
			_, ioErr := io.Copy(out, file)
			if ioErr == nil {
				res := map[string]interface{}{
					"gid":        gid,
					"image_path": imagePath,
					"image_name": fileName,
				}
				context.JSON(http.StatusOK, gin.H{
					"message": "Upload successfully",
					"data":    res,
					"code":    201,
				})
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{
					"error":  "Upload failed",
					"detail": ioErr.Error(),
					"code":   501,
				})
			}
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Upload failed",
				"detail": err.Error(),
				"code":   502,
			})
		}
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Upload failed",
			"detail": err.Error(),
			"code":   503,
		})
	}

}
