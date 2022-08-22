package model

type AdminPermission struct {
	Id           int64 `gorm:"primaryKey"`
	AdminId      int64
	Admin        Admin `gorm:"foreignKey:AdminId"`
	PermissionId int64
	Permission   Permission `gorm:"foreignKey:PermissionId"`
}
