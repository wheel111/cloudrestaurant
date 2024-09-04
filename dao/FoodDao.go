package dao

import (
	"cloudrestaurant/model"
	"cloudrestaurant/tool"
)

type FoodDao struct {
	*tool.Orm
}

func NewFoodDao() *FoodDao {
	return &FoodDao{tool.DbEngine}
}

// 根据商户id查询所有食品数据
func (fd *FoodDao) QueryFoods(shop_id int64) ([]model.Foods, error) {
	var foods []model.Foods
	err := fd.Orm.Where("shop_id = ?", shop_id).Find(&foods)
	if err != nil {
		return nil, err
	}
	return foods, nil
}
