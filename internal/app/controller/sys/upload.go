package sys

import (
	"bytes"
	"fast-go/pkg/e"
	"fast-go/pkg/logging"
	"fast-go/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)
type Upload struct{}

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")

	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}else{
		var buff = make([]byte,image.Size)
		file.Read(buff)
		url := upload.Upload(image.Filename, "", bytes.NewReader(buff))
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": url,
		})
	}
}

