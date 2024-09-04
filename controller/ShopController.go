package controller

import (
	"cloudrestaurant/service"
	"cloudrestaurant/tool"
	"github.com/gin-gonic/gin"
)

type ShopController struct{}

func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/shops", sc.GetShopList)
	engine.GET("/api/search_shops", sc.SearchShop)
}

// 关键词搜索商户信息
func (sc *ShopController) SearchShop(c *gin.Context) {
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")
	keyword := c.Query("keyword")
	if keyword == "" {
		tool.Fail(c, "查询错误")
		return
	}
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34" //北京
		latitude = "40.34"
	}
	//执行关键词搜索
	shopService := service.ShopService{}
	Searchshops := shopService.SearchShops(longitude, latitude, keyword)
	if len(Searchshops) != 0 {
		tool.Success(c, Searchshops)
	}
	tool.Fail(c, "查找关键字失败")
}

//获取商铺列表

func (sc *ShopController) GetShopList(c *gin.Context) {
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34" //北京
		latitude = "40.34"
	}
	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude)
	if len(shops) == 0 {
		tool.Fail(c, "未获取商铺信息")
		return
	}
	for _, shop := range shops {
		shopServices := shopService.GetService(shop.Id)
		if len(shopServices) == 0 {
			shop.Support = nil
			tool.Success(c, shop)
		} else {
			shop.Support = shopServices
			tool.Success(c, shop)
		}
	}
}
