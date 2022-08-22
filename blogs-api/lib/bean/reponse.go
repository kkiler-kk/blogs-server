package bean

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Response = &response{}

type response struct {
	Success bool        // 是否返回成功 true = 200 false均为失败
	Code    int         // 返回错误码
	Data    interface{} // 返回数据
	msg     string      // 返回数据
}

func (r *response) ResultSuc(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"msg":     "success",
		"data":    data,
		"success": true,
	})
}

func (r *response) ResultFail(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"code":    code,
		"msg":     msg,
		"success": false,
		"data":    nil,
	})
}
