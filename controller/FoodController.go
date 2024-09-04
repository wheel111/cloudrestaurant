package controller

import (
	"cloudrestaurant/service"
	"cloudrestaurant/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

type FoodController struct {
}

func (fc *FoodController) Router(engine *gin.Engine) {
	engine.GET("/api/foods", fc.GetFoods)
}

// 获取某个商户所提供的食品
func (fc *FoodController) GetFoods(c *gin.Context) {
	shopId, exist := c.GetQuery("shop_id")
	if !exist {
		tool.Fail(c, "请求参数错误，请重试")
		return
	}
	//实例化FoodService
	id, err := strconv.Atoi(shopId)
	if err != nil {
		tool.Fail(c, "请求参数错误，请重试")
		return
	}
	foodService := service.NewFoodService()
	foods := foodService.GetFoods(int64(id))
	if len(foods) == 0 {
		tool.Fail(c, "未查询到相关数据")
		return
	}
	tool.Success(c, foods)
}
