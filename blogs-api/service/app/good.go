package app

import (
	"errors"
	"git.oa00.com/go/mysql"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/model"
	"kkblogs/blogs-api/model/reponse"
)

var GoodLogic = &goodLogic{}

type goodLogic struct {
}

// GetArticlesGoodCount @Title 根据文章id查询文章收藏数量
func (*goodLogic) GetArticlesGoodCount(articles int64) int64 {
	var goodCount int64
	if mysql.Db.Model(&model.ArticleGood{}).Where("article_id = ?", articles).Count(&goodCount).Error != nil {
		return 0
	}
	return goodCount
}

// GetGoodByUserId @Title 根据用户id查询他的收藏
func (l *goodLogic) GetGoodByUserId(userId int64) (result []reponse.Articles, err error) {
	var goodList []model.ArticleGood
	if mysql.Db.Where("user_id = ?", userId).Find(&goodList).RowsAffected <= 0 {
		return nil, errors.New("暂时没有收藏哦")
	}
	var articleIds []int
	for _, good := range goodList {
		articleIds = append(articleIds, good.ArticleId)
	}
	result, _, err = ActivityLogic.ListActivity(ArgsArticlesReq{}, bean.Page{}, Search{}, articleIds)
	if err != nil {
		return nil, err
	}
	return
}

// IsGood @Title 根据文章id用户id 查看是否收藏了此文章 true 未收藏 false 为已收藏
func (*goodLogic) IsGood(articleId int, userId int) bool {
	var good model.ArticleGood
	if mysql.Db.Model(&good).Where("article_id = ? and user_id = ?", articleId, userId).First(&good).RowsAffected <= 0 {
		return true
	}
	return false
}
