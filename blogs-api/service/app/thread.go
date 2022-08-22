package app

import (
	"fmt"
	"git.oa00.com/go/mysql"
	"git.oa00.com/go/redis"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/model"
	"log"
	"sync"
)

var ThreadService = &threadService{}

type threadService struct {
	onec sync.Once
}

// UpdateArticleViewCount @Title 修改阅读量
func (t *threadService) UpdateArticleViewCount(article model.Article) {
	go func() {
		key := fmt.Sprintf("%s%d", bean.ViewCommentCount, article.Id)
		get := redis.Redis.Get(key)
		if get == nil {
			if mysql.Db.Select("id, view_counts").First(&article).Error != nil {
				log.Println("更新失败")
				return
			}
			redis.Redis.Set(key, article.ViewCounts, 60*12*100000)
		}
		redis.Redis.Incr(key)
	}()
}

// UpdateCommentViewCount @Title 评论加一
func (t *threadService) UpdateCommentViewCount(articleId int64) {
	go func() {
		key := fmt.Sprintf("%s%d", bean.ViewCommentCount, articleId)
		get := redis.Redis.Get(key)
		if get == nil {
			var article model.Article
			if mysql.Db.Select("comment_counts").Where("id = ?", articleId).First(&article).Error != nil {
				log.Println("更新失败")
				return
			}
			redis.Redis.Set(key, article.CommentCounts, 60*12*100000)
		}
		redis.Redis.Incr(key)
	}()
}
