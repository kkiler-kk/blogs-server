package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/service/app"
	"strconv"
)

type Tag struct {
}

// Hot @Title 最热标签
func (*Tag) Hot(c *gin.Context) {
	limit := bean.LIMIT
	result, err := app.TagsLogic.ListHot(limit)
	if err != nil {
		bean.Response.ResultFail(c, 10002, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

// ListTags @Title 查询所有标签
func (t *Tag) ListTags(c *gin.Context) {
	result, err := app.TagsLogic.ListTags()
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

// Detail @Title 查询所有标签
func (t *Tag) Detail(c *gin.Context) {
	result, err := app.TagsLogic.FindAllDetail()
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

// DetailId @Title 查询标签根据id
func (t *Tag) DetailId(c *gin.Context) {
	id := c.Param("id")
	// string到int64
	tagId, _ := strconv.ParseInt(id, 10, 64)
	result, err := app.TagsLogic.DetailId(tagId)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}
