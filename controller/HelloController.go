package controller

import "github.com/gin-gonic/gin"

type HelloController struct {
}

func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/Hello", hello.Hello)
}

func (hello *HelloController) Hello(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"message": "Hello cloudrestaurant",
	})
}
