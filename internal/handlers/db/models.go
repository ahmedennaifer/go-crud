package db

type Book struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Content string
	Author  string
	UserId  uint `gorm:"foreignKey"`
}

type User struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Password string
}
