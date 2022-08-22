package app

import (
	"errors"
	"git.oa00.com/go/mysql"
	"kkblogs/blogs-api/model"
	"kkblogs/blogs-api/model/reponse"
)

var CategoryLogic = &categoryLogic{}

type categoryLogic struct {
}

func (*categoryLogic) ListCategory() (result []reponse.ListCategory, err error) {
	var categorys []model.Category
	if mysql.Db.Find(&categorys).Error != nil {
		return result, errors.New("暂无分类")
	}
	for _, category := range categorys {
		result = append(result, reponse.ListCategory{
			Id:           category.Id,
			Avatar:       category.Avatar,
			CategoryName: category.CategoryName,
		})
	}
	return
}

// FindAllDetail @Title 查询所有分类标签
func (l *categoryLogic) FindAllDetail() (result []reponse.ListCategory, err error) {
	var categorys []model.Category
	if mysql.Db.Find(&categorys).Error != nil {
		return result, errors.New("暂无分类")
	}
	for _, category := range categorys {
		result = append(result, reponse.ListCategory{
			Id:           category.Id,
			Avatar:       category.Avatar,
			CategoryName: category.CategoryName,
			Description:  category.Description,
		})
	}
	return
}

// DetailId @Title 查询分类 根据id
func (l *categoryLogic) DetailId(id int64) (result reponse.ListCategory, err error) {
	var category = model.Category{Id: id}
	if mysql.Db.First(&category).Error != nil {
		return result, errors.New("暂无分类")
	}
	result = reponse.ListCategory{
		Id:           category.Id,
		Avatar:       category.Avatar,
		CategoryName: category.CategoryName,
		Description:  category.Description,
	}
	return result, err
}
