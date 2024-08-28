package tool

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)

// session 中间件初始化
func InitSession(engine *gin.Engine) {
	config := GetConfig().RedisConfig
	store, err := redis.NewStore(10, "tcp", config.Addr+":"+config.Port, config.Password, []byte("secret"))
	if err != nil {
		log.Fatal(err.Error())
	}
	engine.Use(sessions.Sessions("mysession", store))
}

//set session

func SetSession(c *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(c)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

//get session

func GetSession(c *gin.Context, key interface{}) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}
