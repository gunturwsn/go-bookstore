package models

import (
	"github.com/gunturwsn/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book represents a book object.
type Book struct {
	gorm.Model
	// Name        string `gorm:""json:"name"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook represent API to Create Book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBooks represent API for get all books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookByID represent API for get book by ID
func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

// DeleteBook represent API for delete book
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
