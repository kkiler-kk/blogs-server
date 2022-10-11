package app

import (
	"errors"
	"git.oa00.com/go/mysql"
	"kkblogs/blogs-api/common"
	"kkblogs/blogs-api/model"
	"kkblogs/blogs-api/model/reponse"
	"time"
)

var UserLogic = &userLogic{}

type userLogic struct {
}

// GetUserById @Title 根据id查询用户信息
func (l userLogic) GetUserById(userId int) (result reponse.UserRep, err error) {
	var user = model.SysUser{Id: int64(userId)}
	if mysql.Db.First(&user).Error != nil {
		return result, errors.New("请重新登录！")
	}
	if user.Signature == "" {
		user.Signature = "Ghetto Love"
	}
	result = reponse.UserRep{
		Id:        user.Id,
		Account:   user.Account,
		NikeName:  user.NickName,
		Avatar:    user.Avatar,
		Signature: user.Signature,
	}
	return
}

// GetUserInfo @Title 获取用户信息
func (l userLogic) GetUserInfo(userId int64) (result reponse.UserInfoRep, err error) {
	var user = model.SysUser{Id: userId}
	if mysql.Db.First(&user).Error != nil {
		return result, errors.New("系统繁忙, 稍后重试")
	}
	result = reponse.UserInfoRep{
		Id:        user.Id,
		Account:   user.Account,
		NikeName:  user.NickName,
		Avatar:    user.Avatar,
		Signature: user.Signature,
		City:      user.City,
		BirthDate: user.BirthDate.UnixMilli(),
		Phone:     user.Phone,
		Email:     user.Email,
		Age:       user.Age,
		Gender:    user.Gender,
	}
	return
}

// ChangePwd @Title 修改密码
func (l userLogic) ChangePwd(userId int64, oldPassword, newPassword string) error {
	var user = model.SysUser{Id: userId}
	if mysql.Db.Select("password, salt").First(&user).Error != nil {
		return errors.New("密码错误")
	}
	oldPassword = common.GetPassword(oldPassword, user.Salt)
	if user.Password != oldPassword {
		return errors.New("密码错误")
	}
	newPassword = common.GetPassword(newPassword, user.Salt)
	if mysql.Db.Model(&user).Update("password", newPassword).RowsAffected <= 0 {
		return errors.New("修改失败")
	}
	return nil
}

type ArgsUserInfoReq struct {
	Password  string
	NikeName  string
	Avatar    string
	Signature string
	City      string
	BirthDate int64
	Phone     string
	Email     string
}

// EditInfo @Title 修改资料
func (l userLogic) EditInfo(userId int64, args ArgsUserInfoReq) error {
	var user = model.SysUser{
		Id: userId,
	}
	if mysql.Db.Model(&user).Where("id = ?", userId).First(&user).Error != nil {
		return errors.New("更新失败")
	}
	var newUser = model.SysUser{
		Id:        userId,
		Avatar:    args.Avatar,
		Signature: args.Signature,
		BirthDate: time.UnixMilli(args.BirthDate),
		City:      args.City,
		Email:     args.Email,
		Phone:     args.Phone,
		NickName:  args.NikeName,
		Password:  common.GetPassword(args.Password, user.Salt),
	}
	if mysql.Db.Model(&newUser).Updates(&newUser).RowsAffected <= 0 {
		return errors.New("更新失败")
	}
	return nil
}

// MyArticle @Title 查询用户文章
func (l userLogic) MyArticle(userId int) (result []reponse.Articles, err error) {
	where := mysql.Db
	var activitys []model.Article
	if userId != 0 {
		where = where.Where("author_id = ?", userId)
	}
	var tags []model.ArticleTag
	var activityIds []int64
	// 查询文章列表
	if mysql.Db.Preload("User").Where(where).Order("weight desc").Order("created_at").Find(&activitys).Error != nil {
		return result, errors.New("系统繁忙")
	}
	// 提取Id
	for _, activity := range activitys {
		activityIds = append(activityIds, activity.Id)
	}
	// 查询标签列表
	if mysql.Db.Preload("Tag").Where("article_id in ?", activityIds).Find(&tags).Error != nil {
		return result, errors.New("系统繁忙")
	}
	// 提取map key：活动id values： tag
	var mapTag = make(map[int64][]reponse.ListTag)
	for _, tag := range tags {
		mapTag[tag.ArticleId] = append(mapTag[tag.ArticleId], reponse.ListTag{
			Id:      tag.Tag.Id,
			TagName: tag.Tag.TagName,
		})
	}
	for _, activity := range activitys {
		result = append(result, reponse.Articles{
			Id:            activity.Id,
			Title:         activity.Title,
			Summary:       activity.Summary,
			CommentCounts: activity.CommentCounts,
			ViewCount:     activity.ViewCounts,
			Weight:        activity.Weight,
			CreateDate:    activity.CreatedAt.UnixNano() / 1e6,
			Author:        activity.User.NickName,
			Tags:          mapTag[activity.Id],
		})
	}
	return
}
