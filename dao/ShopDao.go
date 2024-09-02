package dao

import (
	"cloudrestaurant/model"
	"cloudrestaurant/tool"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() *ShopDao {
	return &ShopDao{tool.DbEngine}
}

const DEFAULT_RANGE = 5

func (shopDao *ShopDao) QueryShop(long, lati float64) []model.Shop {
	var shops []model.Shop
	err := shopDao.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ?", long-DEFAULT_RANGE, long+DEFAULT_RANGE, lati-DEFAULT_RANGE, lati+DEFAULT_RANGE).Find(&shops)
	if err != nil {
		return nil
	}
	return shops
}
