package model

type ArticleTag struct {
	Id        int64 `gorm:"primaryKey"`
	ArticleId int64
	Article   Article `gorm:"foreignKey:ArticleId"`
	TagId     int64
	Tag       Tag `gorm:"foreignKey:TagId"`
}
