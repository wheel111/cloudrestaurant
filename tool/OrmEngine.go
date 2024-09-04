package tool

import (
	"cloudrestaurant/model"
	"fmt"
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
	err = engine.Sync2(new(model.SmsCode), new(model.Member), new(model.FoodCategory), new(model.Shop), new(model.Service), new(model.ShopService), new(model.Foods))
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	// 插入初始化shops数据
	InitShopData()
	// 插入初始化foods数据
	InitFoodsData()
	return orm, nil
}

// 向shop表中初始化数据
func InitShopData() {
	shops := []model.Shop{}
	//事务
	session := DbEngine.NewSession()
	defer session.Close()
	//事务操作： 事务开始， 执行操作（回滚），提交事务
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, shop := range shops {
		_, err := session.Insert(&shop)
		if err != nil {
			session.Rollback()
			return
		}
		err = session.Commit()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

// 初始化Foods表格
func InitFoodsData() {
	foods := []model.Foods{}
	//事务
	session := DbEngine.NewSession()
	defer session.Close()
	//事务操作： 事务开始， 执行操作（回滚），提交事务
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, food := range foods {
		_, err := session.Insert(&food)
		if err != nil {
			session.Rollback()
			return
		}
		err = session.Commit()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
