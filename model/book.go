package model

//定义书籍表结构
type Book struct {
	Id    int64  `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null" binding:"required"`
	Info  string `json:"info" gorm:"not null" binding:"required"`
	Users []User `gorm:"many2many:book_users"` //关联多对多，表名为book_users
}
