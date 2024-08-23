package controller

import "github.com/gin-gonic/gin"

type HelloController struct {
}

// hello接口
func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/Hello", hello.Hello)
}

// hello接口响应成功返回json值
func (hello *HelloController) Hello(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"message": "Hello cloudrestaurant",
	})
}
