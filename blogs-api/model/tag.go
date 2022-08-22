package model

type Tag struct {
	Id      int64 `gorm:"primaryKey"`
	Avatar  string
	TagName string
}
