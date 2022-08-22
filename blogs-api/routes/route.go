package routes

import (
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/routes/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	// 跨域
	router.Use(middleware.Cors)
	// 404
	router.NoRoute(func(c *gin.Context) {
		bean.Response.ResultFail(c, 404, "接口不存在")
	})
	AppRoute(router.Group("app"))
	ManageRoute(router.Group("manage"))
}
