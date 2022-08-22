package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/lib/qiniuyun"
	"log"
	"strings"
)

type Upload struct {
}

// ImagesUpload @Title 文件上传
func (u Upload) ImagesUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	if file == nil {
		bean.Response.ResultFail(c, -101, "请上传正确的图片地址")
		return
	}
	last := strings.Split(file.Filename, ".")
	lastStr := last[1]
	last[1] = common.RandStr(6)
	fileName := fmt.Sprintf("%s%s%s%s", last[0], last[1], ".", lastStr)
	log.Println(fileName)
	_, url := qiniuyun.UploadToQiNiu(file, fileName)
	bean.Response.ResultSuc(c, url)
}
