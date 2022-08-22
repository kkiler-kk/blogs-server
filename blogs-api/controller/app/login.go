package app

import (
	"fmt"
	"git.oa00.com/go/redis"
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/lib/jwt"
	"kkblogs/blogs-api/service/app"
	"time"
)

type Login struct {
}
type argsLoginReq struct {
	Account  string `binding:"required" json:"account"`
	Password string `binding:"required,min=6" label:"密码"`
}

// Login @Title 登录
func (*Login) Login(c *gin.Context) {
	args := argsLoginReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	// 设置redis 防止暴力破解
	key := fmt.Sprintf("%s%s", bean.BRUTE, args.Account)
	count, _ := redis.Redis.Get(key).Int()
	if count >= 5 {
		bean.Response.ResultFail(c, -103, "请求次数过多 请稍后重试吧")
		return
	}
	token, err := app.LoginLogic.Login(args.Account, args.Password)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	redis.Redis.Incr(key)
	redis.Redis.Expire(key, time.Hour*12)
	bean.Response.ResultSuc(c, token)
}

// Logout @Title 退出登录
func (l *Login) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	getTokenKey := fmt.Sprintf("%s%d", bean.TOKEN, parseToken.UserId)
	refreshTokenKey := fmt.Sprintf("%s%d", bean.RefreshToken, parseToken.UserId)
	app.LoginLogic.Logout(getTokenKey, refreshTokenKey)
	bean.Response.ResultSuc(c, nil)
}
