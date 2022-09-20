package model

type ArticleLike struct {
	LikeId    int64 `gorm:"primaryKey"`
	ArticleId int
	Article   Article `gorm:"foreignKey:ArticleId"`
	UserId    int64
	User      SysUser `gorm:"foreignKey:UserId"`
}
