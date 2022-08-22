package model

type Permission struct {
	Id          int64 `gorm:"primaryKey"`
	Name        string
	Path        string
	Description string
}
