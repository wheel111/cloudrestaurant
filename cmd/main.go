package main

import (
	"cloudrestaurant/controller"
	"cloudrestaurant/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	// 注册路由
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

// 注册路由
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
