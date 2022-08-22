package dao

import (
	"git.oa00.com/go/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"kkblogs/blogs-api/model"
)

var CommentsMapper = &commentsMapper{}

type commentsMapper struct {
}

// CreateComment @Title 创建评论
func (receiver commentsMapper) CreateComment(comment model.Comment) error {
	return mysql.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(&comment).RowsAffected <= 0 {
			return errors.New("评论失败")
		}
		return nil
	})
}
