package models

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
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

func GetBook(w http.ResponseWriter, r *http.Request) (*Book, error){

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		fmt.Fprint(w, "Error: ", err)
	}

	book := &Book{
		Id: id,
	}

	err = db.Model(book).WherePK().Select()

	return book, err
}

func UpdateBook(formData url.Values, w http.ResponseWriter, r *http.Request) (*Book, error) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		fmt.Print("Error: ", err)
	}

	book := &Book{
		Id: id,
	}

	err = db.Model(book).WherePK().Select()

	if err != nil {
		panic(err)
	}

	book.Isbn = formData.Get("book_isbn")
	book.Name = formData.Get("book_name")
	book.Author = formData.Get("book_author")
	book.Pages = formData.Get("book_pages")
	book.Year = formData.Get("book_year")

	_, err = db.Model(book).WherePK().Update()

	if err != nil {
		panic(err)
	}

	err = db.Model(book).WherePK().Select()

	if err != nil {
		panic(err)
	}

	return book, err
}

func DeleteBook(w http.ResponseWriter, r *http.Request) bool {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		fmt.Print("Error: ", err)
	}

	book := &Book {
		Id: id,
	}

	_, err = db.Model(book).WherePK().Delete()

	if err != nil {
		return false
	}

	return true
}

func CloseDb() error{
	return db.Close()
}