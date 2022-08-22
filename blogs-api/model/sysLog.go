package model

import "time"

type SysLog struct {
	Id        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	Ip        string
	Method    string
	Module    string
	NickName  string
	Operation string
	Params    string
	Time      time.Time
	UserId    int64
}
