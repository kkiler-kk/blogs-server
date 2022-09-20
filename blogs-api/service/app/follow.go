package app

import (
	"errors"
	"git.oa00.com/go/mysql"
	"gorm.io/gorm"
	"kkblogs/blogs-api/model"
	"kkblogs/blogs-api/model/reponse"
)

var FollowLogic = &followLogic{}

type followLogic struct {
}

//GetUserFollow @Title 根据id 获取用户粉丝 关注 发帖数
func (l followLogic) GetUserFollow(userId int64) (result reponse.FollowFansRep, err error) {
	if mysql.Db.Raw("SELECT ( SELECT count( follow_id ) FROM `tab_user_follow` WHERE follow_id = ? ) fans_count, ( SELECT count( user_id ) FROM `tab_user_follow` WHERE user_id = ? ) follow_count, (select count(id) from `tab_article` where author_id = ? and deleted_at IS NULL) articles_count FROM `tab_user_follow` GROUP BY follow_count, fans_count, articles_count", userId, userId, userId).Scan(&result).Error != nil {
		return result, errors.New("系统繁忙, 稍后重试")
	}
	return
}

// FollowAdd @Title 关注
func (l followLogic) FollowAdd(followId, fansId int64) error {
	var follow = model.UserFollow{FollowId: int(followId), UserId: int(fansId)}
	if mysql.Db.Where("follow_id = ? and user_id = ?", followId, fansId).First(&follow).RowsAffected > 0 {
		return errors.New("已经关注了")
	}
	return mysql.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(&follow).RowsAffected <= 0 {
			return errors.New("稍后重试")
		}
		return nil
	})
}

// IsFollow @Title 是否关注
func (l followLogic) IsFollow(followId, fansId int) (result bool) {
	var userFollow = model.UserFollow{FollowId: followId, UserId: fansId}
	if mysql.Db.Where("follow_id = ? and user_id = ?", followId, fansId).First(&userFollow).RowsAffected > 0 {
		return true
	}
	return false
}

// MyFollow @Title 查看用戶关注
func (l followLogic) MyFollow(userId int64) (result []reponse.UserRep, err error) {
	var userFollow []model.UserFollow
	if mysql.Db.Preload("FollowUser").Model(&userFollow).Where("user_id = ?", userId).Find(&userFollow).RowsAffected <= 0 {
		return result, errors.New("暂时没有关注")
	}
	for _, follow := range userFollow {
		result = append(result, reponse.UserRep{
			Id:        int64(follow.FollowId),
			NikeName:  follow.FollowUser.NickName,
			Avatar:    follow.FollowUser.Avatar,
			Signature: follow.FollowUser.Signature,
		})
	}
	return
}

// MyFans @Title 查看关注
func (l followLogic) MyFans(userId int64) (result []reponse.UserRep, err error) {
	var userFollow []model.UserFollow
	if mysql.Db.Preload("UserFans").Model(&userFollow).Where("follow_id = ?", userId).Find(&userFollow).RowsAffected <= 0 {
		return result, errors.New("暂时没有关注")
	}
	for _, follow := range userFollow {
		result = append(result, reponse.UserRep{
			Id:        int64(follow.UserId),
			NikeName:  follow.UserFans.NickName,
			Avatar:    follow.UserFans.Avatar,
			Signature: follow.UserFans.Signature,
		})
	}
	return
}

// FollowRemove @Title 取消关注
func (l followLogic) FollowRemove(followId int64, fansId int64) error {
	var follow = model.UserFollow{FollowId: int(followId), UserId: int(fansId)}
	return mysql.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("follow_id = ? and user_id = ?", followId, fansId).Delete(&follow).RowsAffected <= 0 {
			return errors.New("稍后重试")
		}
		return nil
	})
}
