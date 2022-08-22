package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func Cors(c *gin.Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")

	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Headers", "Content-Type, Device-Type,Authorization, Manage-Token, User-Token, X-Requested-With, oauth-token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT, WS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Set("content-type", "application/json") //设置返回格式是json

	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}
