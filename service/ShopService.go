package service

import (
	"cloudrestaurant/dao"
	"cloudrestaurant/model"
	"strconv"
)

type ShopService struct {
}

// 根据shopId获取商家服务
func (sc *ShopService) GetService(shopId int64) []model.Service {
	shopDao := dao.NewShopDao()
	return shopDao.QueryServiceByShopId(shopId)
}

// 查询商铺列表数据
func (sc *ShopService) ShopList(long, lat string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}
	shopDao := dao.NewShopDao()
	return shopDao.QueryShop(longitude, latitude, "")
}

// 根据关键字查询对应商家信息

func (sc *ShopService) SearchShops(long, lat, keyword string) []model.Shop {
	shopDao := dao.NewShopDao()
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}
	return shopDao.QueryShop(longitude, latitude, keyword)
}
