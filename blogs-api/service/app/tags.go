package app

import (
	"errors"
	"git.oa00.com/go/mysql"
	"kkblogs/blogs-api/model"
	"kkblogs/blogs-api/model/reponse"
)

var TagsLogic = &tagsLogic{}

type tagsLogic struct {
}

// ListHot @Title 返回最热标签
func (l tagsLogic) ListHot(limit int) (result []reponse.ListTag, err error) {
	var articleTag model.ArticleTag
	var tagIds []int
	if mysql.Db.Model(&articleTag).Select("tag_id").Group("tag_id").Order("count(*) desc").Limit(limit).Find(&tagIds).Error != nil {
		return result, errors.New("系统繁忙, 稍后重试")
	}
	var tagList []model.Tag
	if mysql.Db.Where("id in ?", tagIds).Find(&tagList).Error != nil {
		return result, errors.New("系统繁忙, 稍后重试")
	}
	for _, tag := range tagList {
		result = append(result, reponse.ListTag{
			Id:      tag.Id,
			TagName: tag.TagName,
		})
	}
	return
}

// ArticlesIdTags @Title 根据文章id查询标签
func (l tagsLogic) ArticlesIdTags(articleId int64) (result []reponse.ListTag, err error) {
	var articleTag model.ArticleTag
	var tagIds []int
	where := mysql.Db
	if articleId != 0 {
		where = where.Where("article_id = ?", articleId)
	}
	if mysql.Db.Model(&articleTag).Where(where).Select("tag_id").Group("tag_id").Order("count(*) desc").Find(&tagIds).Error != nil {
		return result, errors.New("系统繁忙, 稍后重试")
	}
	var tagList []model.Tag
	if mysql.Db.Where("id in ?", tagIds).Find(&tagList).Error != nil {
		return result, errors.New("系统繁忙, 稍后重试")
	}
	for _, tag := range tagList {
		result = append(result, reponse.ListTag{
			Id:      tag.Id,
			TagName: tag.TagName,
		})
	}
	return
}

// ListTags @Title 查询所有标签
func (l tagsLogic) ListTags() (result []reponse.ListTag, err error) {
	var tags []model.Tag
	if mysql.Db.Find(&tags).Error != nil {
		return result, errors.New("暂无分类")
	}
	for _, tag := range tags {
		result = append(result, reponse.ListTag{
			Id:      tag.Id,
			TagName: tag.TagName,
		})
	}
	return
}

// FindAllDetail @Title 查询所有标签
func (l tagsLogic) FindAllDetail() (result []reponse.ListTag, err error) {
	var tags []model.Tag
	if mysql.Db.Find(&tags).Error != nil {
		return result, errors.New("暂无分类")
	}
	for _, tag := range tags {
		result = append(result, reponse.ListTag{
			Id:      tag.Id,
			TagName: tag.TagName,
			Avatar:  tag.Avatar,
		})
	}
	return
}

// DetailId @Title 查询标签根据id
func (l tagsLogic) DetailId(id int64) (result reponse.ListTag, err error) {
	var tag = model.Tag{Id: id}
	if mysql.Db.First(&tag).Error != nil {
		return result, errors.New("暂无分类")
	}
	result = reponse.ListTag{
		Id:      tag.Id,
		TagName: tag.TagName,
		Avatar:  tag.Avatar,
	}
	return
}
