package tool

import (
	"cloudrestaurant/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// 创建使用数据库框架结构体
type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

// 创建数据库初始化
func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database
	// 创建连接
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" +
		database.Dbname + "?charset=" + database.Charset
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return nil, err
	}
	// 测试连接
	engine.ShowSQL(true)
	// 映射创建表结构
	err = engine.Sync2(new(model.SmsCode), new(model.Member), new(model.FoodCategory), new(model.Shop))
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
