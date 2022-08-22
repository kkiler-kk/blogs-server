package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/model/reponse"
	"kkblogs/blogs-api/service/app"
	"strconv"
)

type Articles struct {
}

// ListArticles @Title 返回首页文章（最新）
func (*Articles) ListArticles(c *gin.Context) {
	args := app.ArgsArticlesReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	result, _, err := app.ActivityLogic.ListActivity(args, args.Page, args.Search)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}

	bean.Response.ResultSuc(c, result)
}

// HotArticles @Title 查询最热文章(Title)
func (a *Articles) HotArticles(c *gin.Context) {
	limit := bean.LIMIT
	result, err := app.ActivityLogic.HNArticles(limit, "view_counts desc")
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

// NewArticles @Title 查询最新文章(Title)
func (a *Articles) NewArticles(c *gin.Context) {
	limit := bean.LIMIT
	result, err := app.ActivityLogic.HNArticles(limit, "created_at desc")
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

type argsListArchivesReq struct {
	UserId int
}

// ListArchives @Title 文章归档
func (a *Articles) ListArchives(c *gin.Context) {
	args := argsListArchivesReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	result, err := app.ActivityLogic.ListArchives(args.UserId)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

// ArticlesView @Title 查看文章详情
func (a *Articles) ArticlesView(c *gin.Context) {
	id := c.Param("id")
	// string到int64
	articlesId, _ := strconv.ParseInt(id, 10, 64)
	result, err := app.ActivityLogic.ArticlesView(articlesId)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

// Publish @Title 发布文章
func (a *Articles) Publish(c *gin.Context) {
	args := app.ArgsArticlesPublishReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	user := bean.Storage.Get().(reponse.LoginUser)
	result, err := app.ActivityLogic.Publish(args, user.Id)
	if err != nil {
		bean.Response.ResultFail(c, -102, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

type argsArticlesSearchReq struct {
	Search string
}

// Search @Title 搜索文章
func (a *Articles) Search(c *gin.Context) {
	args := argsArticlesSearchReq{}
	if err := c.ShouldBind(&args); err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	result, err := app.ActivityLogic.Search(args.Search)
	if err != nil {
		return
	}
	bean.Response.ResultSuc(c, result)
}
