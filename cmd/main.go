package main

import (
	"cloudrestaurant/controller"
	"cloudrestaurant/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
}
