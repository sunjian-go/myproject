package router

import "github.com/gin-gonic/gin"

//一个测试路由
func LoadTestRouter(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "test")
	})
}
