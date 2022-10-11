package router

import (
	"bookmanager/controller"
	"github.com/gin-gonic/gin"
)

//主操作路由
func RegisterRouter(r *gin.Engine) {
	r.POST("/register", controller.AuthHandlerFunc)      //用户注册
	r.POST("/login", controller.LoginFcun)               //用户登录
	r.POST("/book", controller.CreateBookFunc)           //增加书籍
	r.GET("/book/:id", controller.GetBookFunc)           //获取书籍信息
	r.PUT("/book/update/:id", controller.UpdateBookFunc) //修改书籍信息
	r.DELETE("/book/delete", controller.DeleteBookFunc)
}
