package app

import (
	"errors"
	"fmt"
	"git.oa00.com/go/mysql"
	"git.oa00.com/go/redis"
	"gorm.io/gorm"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/lib/jwt"
	"kkblogs/blogs-api/model"
	"time"
)

var RegisterLogic = &registerLogic{}

type registerLogic struct {
}

// Register @Title 注册
func (*registerLogic) Register(account, name, password string) (token string, err error) {
	salt := common.RandStr(6)
	password = common.GetPassword(password, salt)
	var user = model.SysUser{Account: account, NickName: name,
		Password: password, LastLogin: time.Now(), Salt: salt,
		Avatar: "require(@/assets/img/default_avatar.png)", BirthDate: time.Now(),
	}
	erro := mysql.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(&user).RowsAffected <= 0 {
			return errors.New("账号已经存在了，请换一个吧")
		}
		return nil
	})
	generateToken, err := jwt.GenerateToken(user.Id, 12)
	refreshTokenToken, err := jwt.GenerateToken(user.Id, 15)
	if err != nil {
		return "", err
	}
	generateTokenKey := fmt.Sprintf("%s%d", bean.TOKEN, user.Id)
	refreshTokenTokenKey := fmt.Sprintf("%s%d", bean.RefreshToken, user.Id)
	redis.Redis.Set(generateTokenKey, generateToken, bean.TokenPast)
	redis.Redis.Set(refreshTokenTokenKey, refreshTokenToken, bean.RefreshTokenKeyPast)
	return generateToken, erro
}
