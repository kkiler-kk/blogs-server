package middleware

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/lib/jwt"
	"kkblogs/blogs-api/service/app"
)

// AppAuth @Title 拦截用户是否登录
func AppAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("Authorization")
	}
	if token == "undefined" {
		bean.Response.ResultFail(c, -201, "用户未登录")
		c.Abort()
		return
	}
	if token == "" {
		bean.Response.ResultFail(c, -201, "用户未登录")
		c.Abort()
		return
	}
	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		bean.Response.ResultFail(c, -201, "token已过期，请重新登录")
		c.Abort()
		return
	}
	result, err := app.LoginLogic.Auth(parseToken.UserId)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		c.Abort()
		return
	}
	bean.Storage.Set(result)
}
