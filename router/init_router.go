package router

import "github.com/gin-gonic/gin"

//初始化路由引擎并传给其他函数
func InitRoute() *gin.Engine {
	r := gin.Default() //创建路由引擎
	LoadTestRouter(r)
	RegisterRouter(r)
	return r
}
