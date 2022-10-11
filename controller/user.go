package controller

import (
	"bookmanager/dao"
	"bookmanager/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//注册用户
func AuthHandlerFunc(c *gin.Context) {
	p := new(model.User)       //定义一个结构体类型的变量并开辟空间
	err := c.ShouldBindJSON(p) //校验获取的前端请求数据是否符合，若符合则绑定到这个结构体
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	dao.DB.Create(&p) //数据校验通过后，创建用户数据，p里面存的是绑定好的用户传入的数据
	c.JSON(200, "auth successfule")
}

//登录验证
func LoginFcun(c *gin.Context) {
	//验证数据是否合规
	p := new(model.User)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(403, gin.H{
			"msg": err.Error(),
		})
	}
	//合规之后开始校验数据
	//1.根据用户名检索数据
	u := model.User{Username: p.Username}                               //创建一个用户结构体，并赋值传入的username
	res := dao.DB.Where("username=?", p.Username).Find(&u).RowsAffected //根据用户传入的username值的结构体进行检索
	if res == 0 {                                                       //结果为零则说明没有该数据
		c.JSON(403, gin.H{
			"msg": "该用户不存在",
		})
	} else {
		//2.根据检索出来的数据与用户发送的数据对比
		if p.Password == u.Password {
			token := uuid.New().String()           //定义一个随机的字符串作为token
			dao.DB.Model(u).Update("token", token) //根据查找出来的结构体信息更新对应字段，”token“是数据库的字段名，token是要传入的数据值
			c.JSON(200, gin.H{
				"msg":   "登陆成功",
				"token": token,
			})
		} else {
			c.JSON(403, gin.H{
				"msg": "登陆失败",
			})
		}
	}
}
