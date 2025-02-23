package models

import (
	"github.com/jinzhu/gorm"
	"github.com/milan-kovac/packages/config"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string 	`json:"author"`
	Publication string `json:"publication"`
}

func init (){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBoook() *Book{
	db.NewRecord(b)
	db.Create(&b)

	return b;
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB){
	var getBok Book 
	db.Where("ID=?",id).Find(&getBok)
	return &getBok, db
}

func DeleteBook(id int64) Book{
	var book Book
	db.Where("ID=?",id).Delete(&book)
	return book
}