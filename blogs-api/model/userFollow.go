package model

import "time"

type UserFollow struct {
	Id         int       `gorm:"primaryKey"` // 主键
	FollowId   int       // 关注人id
	FollowUser SysUser   `gorm:"foreignKey:FollowId"`
	UserId     int       //用户id
	UserFans   SysUser   `gorm:"foreignKey:UserId"`
	CreatedAt  time.Time // 关注时间
}
