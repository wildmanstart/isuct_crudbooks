package models

import (
	"github.com/go-pg/pg/v10"
	"net/url"
)
// Модель книги
type Book struct {
	Id int64
	Isbn string
	Name string
	Author string
	Pages string
	Year string
	Addedon string
}

// TODO::Конфиг вынести в отдельный файл
var db = pg.Connect(&pg.Options {
	Database: "isuct_base",
	User: "v-drondin",
	Password: "PASS_word",
})

func GetBooks() []Book {
	var books []Book

	err := db.Model(&books).Select()

	if err != nil {
		panic(err)
	}

	return books
}

// TODO::На будущее необходимо изучить дргуие методы передачи, кроме переменной
func CreateBook(formData url.Values) (*Book, error){
	book := &Book{
		Isbn: formData.Get("isbn"),
		Name: formData.Get("name"),
		Author: formData.Get("author"),
		Pages: formData.Get("pages"),
		Year: formData.Get("year"),
	}

	_, err := db.Model(book).Insert()
	if err != nil {
		panic(err)
	}

	return book, nil
}

func CloseDb() error{
	return db.Close()
}