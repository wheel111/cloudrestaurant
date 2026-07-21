package tool

import (
	"blog/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 创建使用数据库框架结构体
type Gorm struct {
	*gorm.DB
}

var DbEngine *Gorm

// 创建数据库初始化
func GormEngine(cfg *Config) (*Gorm, error) {
	database := cfg.Database
	// 创建连接
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" +
		database.Dbname + "?charset=" + database.Charset + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}
	// 映射创建表结构
	err = db.AutoMigrate(&model.User{}, &model.Comment{}, &model.Category{}, &model.Post{}, &model.Tag{})
	if err != nil {
		return nil, err
	}
	gormDB := new(Gorm)
	gormDB.DB = db
	DbEngine = gormDB
	fmt.Println("数据库成功连接")
	return gormDB, nil
}
