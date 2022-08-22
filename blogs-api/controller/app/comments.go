package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/model/reponse"
	"kkblogs/blogs-api/service/app"
	"log"
	"strconv"
)

type Comments struct {
}

// ArticlesCommentId  @Title 根据文章id 查看文章详情评论
func (*Comments) ArticlesCommentId(c *gin.Context) {
	id := c.Param("id")
	// string到int64
	articlesId, _ := strconv.ParseInt(id, 10, 64)
	result, err := app.CommentsLogic.ArticlesCommentId(articlesId)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

type argsCreateCommentReq struct {
	ArticleId int64  `binding:"required"`
	Content   string `binding:"required"`
	Parent    int64
	ToUserId  int64
}

// CreateComment @Title 创建评论
func (*Comments) CreateComment(c *gin.Context) {
	args := argsCreateCommentReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	log.Println(args)
	user := bean.Storage.Get().(reponse.LoginUser)
	err := app.CommentsLogic.CreateComment(args.ArticleId, args.Content, args.Parent, args.ToUserId, user.Id)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, "评论成功")
}
