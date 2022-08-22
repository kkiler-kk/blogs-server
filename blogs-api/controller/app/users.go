package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/lib/jwt"
	"kkblogs/blogs-api/model/reponse"
	"kkblogs/blogs-api/service/app"
	"strconv"
)

type Users struct {
}

// CurrentUser @Title 解析token
func (*Users) CurrentUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("Authorization")
	}
	if token == "" {
		bean.Response.ResultFail(c, -101, "用户未登录")
		c.Abort()
		return
	}
	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		bean.Response.ResultFail(c, -102, "token已过期，请重新登录")
		c.Abort()
		return
	}
	auth, err := app.LoginLogic.Auth(parseToken.UserId)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		c.Abort()
		return
	}
	bean.Response.ResultSuc(c, auth)
}

// GetUserById @Title 获取用户信息
func (u *Users) GetUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	response, _ := app.UserLogic.GetUserById(userId)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		c.Abort()
		return
	}
	bean.Response.ResultSuc(c, response)
}

// GetUserInfo @Title 获取用户基本信息
func (u *Users) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	info, err := app.UserLogic.GetUserInfo(userId)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, info)
}

type argsUserChangePwdReq struct {
	OldPassword string `binding:"required,min=6" label:"旧密码"`
	NewPassword string `binding:"required,min=6" label:"新密码"`
}

// ChangePwd @Title 修改密码
func (*Users) ChangePwd(c *gin.Context) {
	args := argsUserChangePwdReq{}
	user := bean.Storage.Get().(reponse.LoginUser)
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	err := app.UserLogic.ChangePwd(user.Id, args.OldPassword, args.NewPassword)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, nil)
}

// EditInfo @Title 修改资料
func (u *Users) EditInfo(c *gin.Context) {
	args := app.ArgsUserInfoReq{}
	user := bean.Storage.Get().(reponse.LoginUser)
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	err := app.UserLogic.EditInfo(user.Id, args)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, nil)
}

// MyArticle @Title 根据用户id 查询文章
func (u *Users) MyArticle(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	article, err := app.UserLogic.MyArticle(userId)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, article)

}
