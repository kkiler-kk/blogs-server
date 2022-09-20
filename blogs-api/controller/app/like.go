package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/service/app"
)

type Like struct {
}
type argsIsLikeReq struct {
	ArticleId int
	UserId    int
}

// IsLike @Title 根据文章id 用户id 查看是否有没有收藏此喜欢
func (l Like) IsLike(c *gin.Context) {
	args := argsIsLikeReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	flag := app.LikeLogic.IsLike(args.ArticleId, args.UserId)
	bean.Response.ResultSuc(c, flag)
}
