package model

import (
	"gorm.io/gorm"
	"time"
)

type SysUser struct {
	Id        int64          `gorm:"primaryKey"` // 主键
	Account   string         // 账号
	Admin     int            // 是否为管理员
	Avatar    string         // 头像
	Signature string         // 个性签名
	BirthDate time.Time      // 出生日期
	Age       int            // 年龄
	Gender    int            // 性别 1 男 2 女 3未知
	City      string         // 所在城市
	CreatedAt time.Time      // 注册时间
	DeletedAt gorm.DeletedAt // 是否删除
	Email     string         // 邮箱
	LastLogin time.Time      // 最后登录时间
	Phone     string         `gorm:"column:mobile_phone_number"` // 手机号
	NickName  string         // 昵称
	Password  string         // 密码
	Salt      string         //加密盐
	Status    uint           // 状态
}
