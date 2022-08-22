package app

import (
	"errors"
	"fmt"
	"git.oa00.com/go/mysql"
	"git.oa00.com/go/redis"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/lib/jwt"
	"kkblogs/blogs-api/model"
	"kkblogs/blogs-api/model/reponse"
)

var LoginLogic = &loginLogic{}

type loginLogic struct {
}

// Login @Title 登录
func (*loginLogic) Login(account, password string) (token string, err error) {
	var user = model.SysUser{Account: account}
	if mysql.Db.Where(&user).First(&user).Error != nil {
		return token, errors.New("账号或密码错误")
	}
	password = common.GetPassword(password, user.Salt)
	if user.Password != password {
		return token, errors.New("账号或密码错误")
	}
	getToken, err := jwt.GenerateToken(user.Id, 12)        // 生成token
	refreshGetToken, err := jwt.GenerateToken(user.Id, 15) // 生成token
	if err != nil {
		return "", err
	}
	getTokenKey := fmt.Sprintf("%s%d", bean.TOKEN, user.Id)
	refreshGetTokenKey := fmt.Sprintf("%s%d", bean.RefreshToken, user.Id)
	redis.Redis.Set(getTokenKey, getToken, bean.TokenPast)
	redis.Redis.Set(refreshGetTokenKey, refreshGetToken, bean.RefreshTokenKeyPast)
	return getToken, nil
}

// Auth @Title  验证登录
func (*loginLogic) Auth(userId int64) (result reponse.LoginUser, err error) {
	var user = model.SysUser{Id: userId}
	if mysql.Db.First(&user).Error != nil {
		return result, errors.New("请重新登录！")
	}
	if user.Signature == "" {
		user.Signature = "Ghetto Love"
	}
	result = reponse.LoginUser{
		Id:        user.Id,
		Account:   user.Account,
		NikeName:  user.NickName,
		Avatar:    user.Avatar,
		Signature: user.Signature,
	}
	return
}

// Logout @Title 退出登录
func (*loginLogic) Logout(getTokenKey, refreshTokenKey string) {
	redis.Redis.Del(getTokenKey)
	redis.Redis.Del(refreshTokenKey)
}
