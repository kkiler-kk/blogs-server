package model

type Body struct {
	Id          int64 `gorm:"primaryKey"`
	Content     string
	ContentHtml string
	ArticleId   int64
	Article     *Article `gorm:"foreignKey:ArticleId"`
}

func (Body) TableName() string {
	return "tab_article_body"
}
