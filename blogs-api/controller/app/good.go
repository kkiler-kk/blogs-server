package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/service/app"
	"strconv"
)

type Good struct {
}

// GetGoodByUserId @Title 根据用户id查询他的收藏
func (t *Good) GetGoodByUserId(c *gin.Context) {
	id := c.Param("id")
	// string到int64
	userId, _ := strconv.ParseInt(id, 10, 64)
	result, err := app.GoodLogic.GetGoodByUserId(userId)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

type argsIsGoodReq struct {
	ArticleId int
	UserId    int
}

// IsGood @Title 根据文章id 用户id 查询是否收藏了此文章
func (t *Good) IsGood(c *gin.Context) {
	args := argsIsGoodReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	flag := app.GoodLogic.IsGood(args.ArticleId, args.UserId)
	bean.Response.ResultSuc(c, flag)
}
