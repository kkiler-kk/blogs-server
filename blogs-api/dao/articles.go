package dao

import (
	"git.oa00.com/go/mysql"
	"github.com/pkg/errors"
	"kkblogs/blogs-api/model"
)

var ActivityMapper = &activityMapper{}

type activityMapper struct {
}

// ArticlesView @Title 查询文章
func (*activityMapper) ArticlesView(articlesId int64) (result model.Article, err error) {
	var articles model.Article
	mysql.Db.Preload("User").Preload("Body").Preload("Category").Where("id = ?", articlesId).First(&articles)
	if articles.Id <= 0 {
		return result, errors.New("文章不存在！")
	}
	return articles, nil
}
