package model

type Book_users struct {
	UserID int64 `gorm:"primaryKey"` //必须定义为主键
	BookID int64 `gorm:"primaryKey"` //必须定义为主键
}
