package model

// Admin 管理员表
type Admin struct {
	Id       int    `gorm:"primaryKey"` // 主键
	UserName string // 用户名
	Password string // 密码
	Salt     string // 加密盐
}
