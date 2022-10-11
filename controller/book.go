package controller

import (
	"bookmanager/dao"
	"bookmanager/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

//增加书籍
func CreateBookFunc(c *gin.Context) {
	//校验传入参数
	p := new(model.Book)                    //记得先开辟空间
	if err := c.ShouldBind(p); err != nil { //校验参数之后，将body里面的数据绑定到结构体
		c.JSON(403, gin.H{
			"msg": "参数不合规",
		})
	} else {
		//参数通过校验开始存入数据库
		if res := dao.DB.Create(&p).RowsAffected; res == 0 {
			c.JSON(200, gin.H{
				"msg": "书籍添加失败",
			})
		} else {
			c.JSON(200, gin.H{
				"msg": "书籍添加成功",
			})
		}
	}
}

//遍历书籍
func GetBookFunc(c *gin.Context) {
	bookIdStr := c.Param("id")                            //获取url中的id值
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)      //将字符串转为10禁止，int64类型的
	book := model.Book{Id: bookId}                        //创建结构体的时候直接赋值
	if res := dao.DB.Find(&book).RowsAffected; res != 0 { //根据传入的ID值的结构体全量查找数据
		c.JSON(200, gin.H{
			"book": book,
		})
	} else {
		c.JSON(403, gin.H{
			"msg": "该书籍未找到",
		})
	}
}

//修改书籍,根据url里面的id修改对应数据
func UpdateBookFunc(c *gin.Context) {
	//获取旧数据的id
	oldId := c.Param("id")                      //获取url中的id
	oldid, _ := strconv.ParseInt(oldId, 10, 64) //将字符串转为10禁止64位的值
	//数据校验
	p := new(model.Book)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(403, gin.H{
			"msg": "数据校验失败",
		})
	} else {
		//数据校验通过后，开始更新数据
		//1.根据id进行更新
		b := model.Book{Id: oldid}                     //定义一个结构体传入旧数据的id
		res := dao.DB.Model(b).Updates(p).RowsAffected //updates可以将结构体内的所有数据对应表结构进行更新
		if res == 0 {
			c.JSON(403, gin.H{
				"msg": "更新失败",
			})
		} else {
			c.JSON(200, gin.H{
				"msg": "更新成功",
			})
		}
	}
}

//删除书籍
func DeleteBookFunc(c *gin.Context) {
	//数据校验
	p := new(model.Book)
	c.ShouldBind(&p) //直接获取数据绑定，不用校验
	if res := dao.DB.Where("id=?", p.Id).Delete(&p).RowsAffected; res == 0 {
		//if res := dao.DB.Model(p).Delete(p).RowsAffected; res == 0 {
		c.JSON(403, gin.H{
			"msg": "删除失败",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "删除成功",
		})
	}
}
