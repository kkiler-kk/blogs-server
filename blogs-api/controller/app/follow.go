package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/model/reponse"
	"kkblogs/blogs-api/service/app"
	"strconv"
)

type Follow struct {
}

// GetUserFollow @Title 获取用户粉丝数 关注数 发帖数
func (f Follow) GetUserFollow(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	result, err := app.FollowLogic.GetUserFollow(userId)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

type argsFollowAddReq struct {
	FollowId int64 `binding:"required"`
}

// AddFollow @Title 关注
func (f Follow) AddFollow(c *gin.Context) {
	user := bean.Storage.Get().(reponse.LoginUser)
	args := argsFollowAddReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	err := app.FollowLogic.FollowAdd(args.FollowId, user.Id)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, nil)
}

// RemoveFollow @Title 取消关注
func (f Follow) RemoveFollow(c *gin.Context) {
	user := bean.Storage.Get().(reponse.LoginUser)
	args := argsFollowAddReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	err := app.FollowLogic.FollowRemove(args.FollowId, user.Id)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, nil)
}

type argsIsFollowReq struct {
	FollowId int `binding:"required"`
	FansId   int `binding:"required"`
}

// IsFollow @Title 博主id 和粉丝id 看是否关注过
func (*Follow) IsFollow(c *gin.Context) {
	args := argsIsFollowReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	result := app.FollowLogic.IsFollow(args.FollowId, args.FansId)
	bean.Response.ResultSuc(c, result)
}

// MyFollow @Title 查询用户关注
func (f Follow) MyFollow(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	result, err := app.FollowLogic.MyFollow(userId)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

// MyFans @Title 查看粉丝
func (f Follow) MyFans(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	result, err := app.FollowLogic.MyFans(userId)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}
