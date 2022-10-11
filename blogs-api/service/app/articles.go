package app

import (
	"errors"
	"fmt"
	"git.oa00.com/go/mysql"
	"git.oa00.com/go/redis"
	"gorm.io/gorm"
	"kkblogs/blogs-api/dao"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/model"
	"kkblogs/blogs-api/model/reponse"
	"strconv"
	"time"
)

var ActivityLogic = &activityLogic{}

type activityLogic struct {
}
type ArgsArticlesReq struct {
	bean.Page
	CategoryId string
	TagId      string
	Year       string
	Month      string
	Search
}
type Search struct {
	UserId int64
}

// ListActivity @Title 获取文章列表
func (*activityLogic) ListActivity(args ArgsArticlesReq, page bean.Page, search Search, articles []int) (result []reponse.Articles, total int64, err error) {
	categoryId, _ := strconv.ParseInt(args.CategoryId, 10, 64)
	tagId, _ := strconv.ParseInt(args.TagId, 10, 64)
	var activitys []model.Article
	where := mysql.Db
	if articles != nil && len(articles) > 0 {
		where = where.Where("id in ?", articles)
	}
	if categoryId != 0 {
		where = where.Where("category_id = ?", categoryId)
	}
	if args.Year != "" {
		where = where.Where("year(created_at) = ?", args.Year)
	}
	if args.Month != "" {
		where = where.Where("month(created_at) = ?", args.Month)
	}
	if search.UserId != 0 {
		where = where.Where("author_id = ?", search.UserId)
	}
	mysql.Db.Model(&activitys).Where(where).Count(&total) // 分页
	if page.HasPage(total) {
		var tags []model.ArticleTag
		var activityIds []int64
		if tagId != 0 { // 根据标签查询文章
			if mysql.Db.Preload("Tag").Where("tag_id = ?", tagId).Find(&tags).Error != nil {
				return result, total, errors.New("系统繁忙")
			}
			for _, activity := range tags {
				activityIds = append(activityIds, activity.ArticleId)
			}
			// 查询文章列表
			if mysql.Db.Preload("User").Where("id in ?", activityIds).Where(where).Offset(page.GetStart()).Limit(page.GetLimit()).Order("weight desc").Order("created_at").Find(&activitys).Error != nil {
				return result, total, errors.New("系统繁忙")
			}
		} else { // 根据文章查询标签
			// 查询文章列表
			if mysql.Db.Preload("User").Where(where).Offset(page.GetStart()).Limit(page.GetLimit()).Order("weight desc").Order("created_at").Find(&activitys).Error != nil {
				return result, total, errors.New("系统繁忙")
			}
			// 提取Id
			for _, activity := range activitys {
				activityIds = append(activityIds, activity.Id)
			}
		}

		// 查询文章喜欢人数
		// 查询标签列表
		if mysql.Db.Preload("Tag").Where("article_id in ?", activityIds).Find(&tags).Error != nil {
			return result, total, errors.New("系统繁忙")
		}
		// 提取map key：活动id values： tag
		var mapTag = make(map[int64][]reponse.ListTag)
		for _, tag := range tags {
			mapTag[tag.ArticleId] = append(mapTag[tag.ArticleId], reponse.ListTag{
				Id:      tag.Tag.Id,
				TagName: tag.Tag.TagName,
			})
		}
		for _, activity := range activitys {
			result = append(result, reponse.Articles{
				Id:            activity.Id,
				Title:         activity.Title,
				Summary:       activity.Summary,
				CommentCounts: activity.CommentCounts,
				LikeCount:     activity.LikeCounts,
				ViewCount:     activity.ViewCounts,
				Weight:        activity.Weight,
				CreateDate:    activity.CreatedAt.Unix(),
				Author:        activity.User.NickName,
				Tags:          mapTag[activity.Id],
			})
		}
	}
	return
}

// HNArticles @Title 查询最热/新文章(Title)
func (l *activityLogic) HNArticles(limit int, desc string) (result []reponse.HNArticles, err error) {
	var articles []model.Article
	if mysql.Db.Select("id, title").Order(desc).Limit(limit).Find(&articles).Error != nil {
		return result, errors.New("系统繁忙， 稍后重试")
	}
	for _, article := range articles {
		result = append(result, reponse.HNArticles{
			Id:    article.Id,
			Title: article.Title,
		})
	}
	return
}

// ListArchives  @Title 文章归档
func (l *activityLogic) ListArchives(userId int) (result []reponse.Archives, err error) {
	var article model.Article
	where := mysql.Db
	if userId != 0 {
		where = where.Where("author_id = ?", userId)
	}
	if mysql.Db.Where(where).Select("YEAR( created_at ) AS `year`,MONTH ( created_at ) AS `month`, count(*) AS count ").Group("year, month").Model(&article).Find(&result).Error != nil {
		return result, errors.New("系统异常， 请稍后重试")
	}
	return
}

// ArticlesView @Title 查看文章详情
func (l *activityLogic) ArticlesView(articlesId int64) (result reponse.ArticlesView, err error) {

	articles, err := dao.ActivityMapper.ArticlesView(articlesId)
	if err != nil {
		return result, err
	}

	// 完成浏览量+1
	ThreadService.UpdateArticleViewCount(articles)
	// 查询标签
	idTags, err := TagsLogic.ArticlesIdTags(articles.Id)
	if err != nil {
		return reponse.ArticlesView{}, err
	}
	viewCountKey := fmt.Sprintf("%s%d", bean.ViewCommentCount, articlesId)
	commentCountKey := fmt.Sprintf("%s%d", bean.CommentCount, articlesId)
	viewCount, _ := redis.Redis.Get(viewCountKey).Int()
	commentCount, _ := redis.Redis.Get(commentCountKey).Int()
	// 查询收藏量
	count := GoodLogic.GetArticlesGoodCount(articlesId)

	result = reponse.ArticlesView{
		Id:            articles.Id,
		Title:         articles.Title,
		Summary:       articles.Summary,
		CommentCounts: uint(commentCount),
		ViewCount:     uint(viewCount),
		Weight:        articles.Weight,
		GoodCount:     uint(count),
		LikeCount:     articles.LikeCounts,
		CreateDate:    articles.CreatedAt.Format("2006-01-02 15:04"),
		Author: reponse.SysUser{
			Id:       articles.User.Id,
			Avatar:   articles.User.Avatar,
			NickName: articles.User.NickName,
		},
		ActivityBody: reponse.ActivityBody{
			Content: articles.Body.Content,
		},
		Tags: idTags,
		Category: reponse.Category{
			Id:           articles.Category.Id,
			Avatar:       articles.Category.Avatar,
			CategoryName: articles.Category.CategoryName,
		},
	}
	return
}

type ArgsArticlesPublishReq struct {
	Title    string `binding:"required"`
	Id       int64
	Body     ArgsArticlesBodyPublishReq `binding:"required"`
	Category ArgsCategoryPublishReq     `binding:"required"`
	Summary  string                     `binding:"required"`
	Tags     []ArgsTagsPublishReq       `binding:"required"`
}
type ArgsArticlesBodyPublishReq struct {
	Content     string
	ContentHtml string
}
type ArgsCategoryPublishReq struct {
	Id           int64
	Avatar       string
	CategoryName string
}
type ArgsTagsPublishReq struct {
	Id int64
}

// Publish @Title 发布文章
func (l *activityLogic) Publish(args ArgsArticlesPublishReq, userId int64) (result reponse.Publish, err error) {
	var article = model.Article{
		Id:    args.Id,
		Title: args.Title,
		Body: model.Body{
			Content:     args.Body.Content,
			ContentHtml: args.Body.ContentHtml,
		},
		CategoryId: int(args.Category.Id),
		Category: model.Category{
			Id:           args.Category.Id,
			Avatar:       args.Category.Avatar,
			CategoryName: args.Category.CategoryName,
		},
		Summary:   args.Summary,
		AuthorId:  userId,
		UpdatedAt: time.Now(),
	}
	var tags []model.ArticleTag

	err = mysql.Db.Transaction(func(tx *gorm.DB) error {
		// 添加
		if article.Id == 0 {
			if tx.Create(&article).Error != nil {
				return errors.New("发布失败, -.-")
			}
		} else { // 修改
			if tx.Omit("created_at").Save(&article).Error != nil {
				return errors.New("发布失败, -.-")
			}
		}

		return nil
	})
	for _, tag := range args.Tags {
		tags = append(tags, model.ArticleTag{ArticleId: article.Id, TagId: tag.Id})
	}
	if mysql.Db.Save(&tags).Error != nil {
		return result, errors.New("发布失败, -.-")
	}
	result = reponse.Publish{Id: article.Id}
	return result, err
}

// Search @Title 搜索文章
func (l *activityLogic) Search(search string) (result []reponse.ActivitySearchRep, err error) {
	if search == "" {
		return result, errors.New("search 不能为空")
	}
	var activitys []model.Article
	if mysql.Db.Select("id, title").Where("title like ?", fmt.Sprintf("%%%s%%", search)).Find(&activitys).RowsAffected <= 0 {
		return result, errors.New("未查询到文章")
	}
	for _, actvity := range activitys {
		result = append(result, reponse.ActivitySearchRep{
			Id:    actvity.Id,
			Title: actvity.Title,
		})
	}
	return
}
