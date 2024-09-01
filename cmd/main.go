package main

import (
	"cloudrestaurant/controller"
	"cloudrestaurant/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	app := gin.Default()
	registerRouter(app)
	// 设置全局跨域访问
	app.Use(Cors())
	// 集成session功能
	tool.InitSession(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

// 注册路由
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
	new(controller.ShopController).Router(router)
}

// 处理跨域请求中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range c.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,"+
				"session")
			c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,"+
				"Access-Control-Allow-Headers,Access-Control-Allow-Methods")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(200, "Options Request!")
		}
		// 继续处理请求
		c.Next()
	}
}
