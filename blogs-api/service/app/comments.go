package app

import (
	"errors"
	"git.oa00.com/go/mysql"
	"kkblogs/blogs-api/dao"
	"kkblogs/blogs-api/model"
	"kkblogs/blogs-api/model/reponse"
)

var CommentsLogic = &commentsLogic{}

type commentsLogic struct {
}

// ArticlesCommentId  @Title 根据文章id 查看文章详情评论
func (l commentsLogic) ArticlesCommentId(articlesId int64) (result []reponse.ListCommentsRep, err error) {
	var comments []model.Comment

	// 查询出所以的父级评论
	if mysql.Db.Preload("User").Model(&comments).
		Where("article_id = ? and level = 1", articlesId).Find(&comments).Error != nil {
		return result, errors.New("评论加载失败")
	}

	// 集合所有父评论id
	var parentId []int64
	for _, comment := range comments {
		parentId = append(parentId, comment.Id)
	}
	var toComments []model.Comment
	if mysql.Db.Preload("User").Preload("ToUser").
		Where("parent_id in ?", parentId).Find(&toComments).Error != nil {
		return result, errors.New("评论加载失败")
	}
	var mapComments = map[int64][]reponse.ListCommentsRep{}
	for _, comment := range toComments {
		mapComments[comment.ParentId] = append(mapComments[comment.ParentId], reponse.ListCommentsRep{
			Id: comment.Id,
			Author: reponse.SysUser{
				Id:       comment.User.Id,
				Avatar:   comment.User.Avatar,
				NickName: comment.User.NickName,
			},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("2006-01-02 15:04"),
			Level:      comment.Level,
			ToUser: reponse.SysUser{
				Id:       comment.ToUser.Id,
				Avatar:   comment.ToUser.Avatar,
				NickName: comment.ToUser.NickName,
			},
		})
	}
	for _, comment := range comments {
		result = append(result, reponse.ListCommentsRep{
			Id: comment.Id,
			Author: reponse.SysUser{
				Id:       comment.User.Id,
				Avatar:   comment.User.Avatar,
				NickName: comment.User.NickName,
			},
			Children:   mapComments[comment.Id],
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("2006-01-02 15:04"),
			Level:      comment.Level,
		})
	}
	return
}

// CreateComment @Title 创建评论
func (l commentsLogic) CreateComment(articleId int64, content string, parent, toUserId, userId int64) error {
	var level uint = 1
	if parent != 0 {
		level = 2
	}
	var comment = model.Comment{ArticleId: articleId, Content: content,
		ParentId: parent, ToUid: toUserId, AuthorId: userId, Level: level}
	err := dao.CommentsMapper.CreateComment(comment)
	if err != nil {
		return err
	}
	ThreadService.UpdateCommentViewCount(comment.ArticleId)
	return nil
}
