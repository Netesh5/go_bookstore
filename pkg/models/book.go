package models

import (
	"book_store/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `gorm:"json:name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

func CreateBook(b *Book) *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(books)
	return books
}

func GetBook(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
