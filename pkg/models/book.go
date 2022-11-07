package models

import (
	"github.com/dragranzer/BMS-MVC/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.LoadEnv()
	config.Connect()
	db = config.GetDB()
	if err := db.Exec("DROP TABLE IF EXISTS books").Error; err != nil {
		panic(err)
	}
	book1 := Book{
		Name:        "Importance of the Time",
		Author:      "Secret",
		Publication: "Umbrella",
	}
	db.AutoMigrate(&Book{})
	db.Create(&book1)
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func (b *Book) GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func (b *Book) GetBookbyId(Id int64) *Book {
	var getBook Book
	db.Where("ID=?", Id).Find(&getBook)
	return &getBook
}

func (b *Book) DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}

func (b *Book) UpdateBook(newBook Book) {
	db.Save(&newBook)
}
