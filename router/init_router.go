package router

import "github.com/gin-gonic/gin"

func InitRoute() *gin.Engine {
	r := gin.Default()
	LoadTestRouter(r)
	RegisterRouter(r)
	return r
}
