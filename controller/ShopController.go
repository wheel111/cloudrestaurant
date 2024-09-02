package controller

import (
	"cloudrestaurant/service"
	"cloudrestaurant/tool"
	"github.com/gin-gonic/gin"
)

type ShopController struct{}

func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/shops", sc.GetShopList)
}

//获取商铺列表

func (sc *ShopController) GetShopList(c *gin.Context) {
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")
	if longitude == "" || latitude == "" {
		longitude = "116.34" //北京
		latitude = "40.34"
	}
	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude)
	if len(shops) != 0 {
		tool.Success(c, shops)
		return
	}
	tool.Fail(c, "未获取商铺信息")
}
