package dao

import (
	"cloudrestaurant/model"
	"cloudrestaurant/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

// 实例化数据库对象
func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.DbEngine}
}

// 从数据库查询所有商品
func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	var categories []model.FoodCategory
	if err := fcd.Engine.Find(&categories); err != nil {
		return nil, err
	}
	return categories, nil
}
