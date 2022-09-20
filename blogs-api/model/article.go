package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	ARTICLE_TOP    = 1 // 动态置顶
	ARTICLE_COMMON = 0 // 评论置顶
)

type Article struct {
	Id            int64     `gorm:"primaryKey"`
	CommentCounts uint      // 评论数量
	CreatedAt     time.Time // 创建时间
	Summary       string    // 简介
	Title         string    // 标题
	ViewCounts    uint      // 浏览数量
	Weight        uint      // 是否置顶
	LikeCounts    uint      // 点赞人数
	AuthorId      int64     // 作者id
	User          SysUser   `gorm:"foreignKey:AuthorId"`
	BodyId        int64     // 内容Id
	Body          Body      `gorm:"foreignKey:BodyId"`
	CategoryId    int
	Category      Category       `gorm:"foreignKey:CategoryId"`
	UpdatedAt     time.Time      // 修改时间
	DeletedAt     gorm.DeletedAt // 软删除
}
