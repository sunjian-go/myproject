package main

import (
	"bookmanager/dao"
	"bookmanager/router"
)

func main() {
	//初始化mysql
	dao.InitMysql()
	//初始化路由
	r := router.InitRoute()
	//启动监听
	r.Run(":9000")

}
