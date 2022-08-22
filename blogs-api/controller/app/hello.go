package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/lib/bean"
)

type Hello struct {
}

func (*Hello) Hello(c *gin.Context) {
	get := c.MustGet("userId")
	bean.Response.ResultSuc(c, get)
	return
}
