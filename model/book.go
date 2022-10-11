package model

type Book struct {
	Id    int64  `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Info  string `json:"info"`
	Users []User `gorm:"many2many:book_users"` //关联多对多，表名为book_users
}

func (Book) TableName() string {
	return "book"
}
