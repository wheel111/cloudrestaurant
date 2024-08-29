package controller

import "github.com/gin-gonic/gin"

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/food_category", fcc.foodCategory)

}

func (fcc *FoodCategoryController) foodCategory(c *gin.Context) {
	//调用service功能获取信息

}
