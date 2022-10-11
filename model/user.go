package model

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`                        //设置别名和主键
	Username string `json:"username" gorm:"not null" binding:"required"` //设置别名、不能为空、请求时候不能为空（shouldBind参数校验）
	Password string `json:"password" gorm:"not null" binding:"required"`
	Token    string `json:"token"`
}

//表名默认添加s,自定义表的名字
func (User) TableName() string {
	return "user"
}
