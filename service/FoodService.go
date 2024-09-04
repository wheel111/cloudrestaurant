package service

import (
	"cloudrestaurant/dao"
	"cloudrestaurant/model"
)

type FoodService struct {
}

func NewFoodService() *FoodService {
	return &FoodService{}
}

// 获取商家食品列表
func (fs *FoodService) GetFoods(shop_id int64) []model.Foods {
	foodDao := dao.NewFoodDao()
	foods, err := foodDao.QueryFoods(shop_id)
	if err != nil {
		return nil
	}
	return foods
}
