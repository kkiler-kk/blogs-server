package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/service/app"
)

type Register struct {
}
type argsRegisterReq struct {
	Account  string `binding:"required,min=4" label:"账号"`
	Password string `binding:"required,min=6" label:"密码"`
	NickName string `binding:"required" label:"昵称"`
	Avatar   string `binding:"required" label:"头像"`
}

// Register @Title 注册
func (r Register) Register(c *gin.Context) {
	args := argsRegisterReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	token, err := app.RegisterLogic.Register(args.Account, args.Avatar, args.NickName, args.Password)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, token)
}
