package dao

import (
	"bookmanager/model"
	"fmt"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "root:123456@tcp(192.168.0.100:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	if err = DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Println("创建表结构失败") //asaasass
	}
}
