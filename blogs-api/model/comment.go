package model

import "time"

type Comment struct {
	Id        int64 `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time
	ArticleId int64
	Article   Article `gorm:"foreignKey:ArticleId"`
	AuthorId  int64
	User      SysUser `gorm:"foreignKey:AuthorId"`
	ParentId  int64
	Comment   *Comment `gorm:"foreignKey:ParentId"`
	ToUid     int64
	ToUser    SysUser `gorm:"foreignKey:ToUid"`
	Level     uint
}
