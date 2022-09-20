package app

import (
	"git.oa00.com/go/mysql"
	"kkblogs/blogs-api/model"
)

var LikeLogic = &likeLogic{}

type likeLogic struct {
}

// IsLike @Title 根据文章id用户id 查看是否喜欢了此文章 true 未收藏 false 为已收藏
func (l likeLogic) IsLike(articleId int, userId int) bool {
	var like model.ArticleLike
	if mysql.Db.Model(&like).Where("article_id = ? and user_id = ?", articleId, userId).First(&like).RowsAffected <= 0 {
		return true
	}
	return false
}
