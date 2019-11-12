package Gorm

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database() {
	// 数据库需要自己手动创建
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/learngo?charset=utf8&parseTime=True&loc=Local")
	// Error
	if err != nil {
		log.Panic("连接数据库不成功", err)
	}

	db.LogMode(true)
	db.SingularTable(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	// 同步数据
	migration()
}
