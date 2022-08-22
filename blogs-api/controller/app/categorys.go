package app

import (
	"github.com/gin-gonic/gin"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/service/app"
	"strconv"
)

type Category struct {
}

// ListCategory @Title 查询所有分类
func (*Category) ListCategory(c *gin.Context) {
	result, err := app.CategoryLogic.ListCategory()
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

// Detail @Title 查询所有文章分类
func (*Category) Detail(c *gin.Context) {
	result, err := app.CategoryLogic.FindAllDetail()
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}

// DetailId @Title 查询分类 根据id
func (*Category) DetailId(c *gin.Context) {
	id := c.Param("id")
	// string到int64
	categoryId, _ := strconv.ParseInt(id, 10, 64)
	result, err := app.CategoryLogic.DetailId(categoryId)
	if err != nil {
		bean.Response.ResultFail(c, -101, common.GetVerErr(err))
		return
	}
	bean.Response.ResultSuc(c, result)
}
