package model

type Category struct {
	Id           int64 `gorm:"primaryKey"`
	Avatar       string
	CategoryName string
	Description  string
}
