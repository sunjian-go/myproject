package router

import (
	"bookmanager/model"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/register", func(c *gin.Context) {
		p := new(model.User)
		err := c.ShouldBindJSON(p)
		if err != nil {
			c.JSON(400, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(200, "auth successfule")
	})

}
