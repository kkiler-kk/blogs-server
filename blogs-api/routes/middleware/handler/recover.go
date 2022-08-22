package handler

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/lib/bean"
	"log"
	"net/http"
	"runtime/debug"
)

// Recover @Title 统一错误格式返回
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			bean.Response.ResultFail(c, http.StatusOK, errorToString(r))
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
