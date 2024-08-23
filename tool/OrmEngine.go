package tool

import (
	"cloudrestaurant/model"
	"github.com/go-xorm/xorm"
)

// 创建使用数据库框架结构体
type Orm struct {
	*xorm.Engine
}

// 创建数据库初始化
func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database
	// 创建短信认证结构体映射成数据库表结构
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" +
		database.Dbname + "?charset=" + database.Charset
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return nil, err
	}
	// 测试连接
	engine.ShowSQL(true)
	// 映射创建表结构
	err = engine.Sync2(new(model.SmsCode))
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	return orm, nil
}
