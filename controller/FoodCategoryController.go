package controller

import (
	"cloudrestaurant/service"
	"cloudrestaurant/tool"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/food_category", fcc.foodCategory)

}

func (fcc *FoodCategoryController) foodCategory(c *gin.Context) {
	//调用service功能获取信息
	foodCategoryService := &service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Fail(c, "食品种类数据获取失败")
		return
	}
	//转换格式
	for _, category := range categories {
		if category.ImageUrl != "" { //图片url拼接
			category.ImageUrl = tool.FileServerAddr() + "/" + category.ImageUrl
		}
		tool.Success(c, categories)
	}
}
