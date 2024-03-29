package upload

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func UploadFile(context *gin.Context) {
	file, header, err := context.Request.FormFile("file")
	if err == nil {
		filename := header.Filename
		out, err := os.Create("static/res/uploadFile/" + filename)
		if err == nil {
			defer func(out *os.File) {
				err := out.Close()
				if err != nil {

				}
			}(out)
			_, err = io.Copy(out, file)
			if err == nil {
				log.Println("上传文件成功")
				res := map[string]interface{}{
					"filePath": "/res/uploadFile/" + filename,
					"fileName": filename,
				}
				context.JSON(http.StatusOK, gin.H{
					"message": "Upload successfully",
					"result":  res,
				})
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{
					"error":  "Upload failed",
					"detail": err.Error(),
				})
			}
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Upload failed",
				"detail": err.Error(),
			})
		}
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Upload failed",
			"detail": err.Error(),
		})
	}
}
