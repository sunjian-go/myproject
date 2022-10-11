package main

import (
	"bookmanager/dao"
	"bookmanager/router"
)

func main() {
	//初始化mysql
	//fmt.Println("mysql db: ", dao.DB)
	dao.InitMysql()
	//初始化路由
	r := router.InitRoute()
	r.Run(":9000")

}
