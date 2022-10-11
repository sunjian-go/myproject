package dao

import (
	"bookmanager/model"
	"fmt"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//初始化mysql操作
func InitMysql() {
	dsn := "root:123456@tcp(192.168.0.100:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{}) //连接数据库
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	if err = DB.AutoMigrate(model.User{}, model.Book{}); err != nil { //创建表结构
		fmt.Println("创建表结构失败")
	}
}
