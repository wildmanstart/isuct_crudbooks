package controllers

import (
	"fmt"
	"net/http"
	"github.com/foolin/goview"
	"main/pkg/models"
)

// Тут мы переопределяем путь до папки с views
var gv = goview.New(goview.Config{
	Root:      "pkg/views", //template root path
	Extension: ".html", //file extension
	Master:    "layouts/master", //master layout file
})

func Home(w http.ResponseWriter, r *http.Request) {
	view := gv.Render(w, http.StatusOK, "home", goview.M {
		"title": "Home page",
	})

	if view != nil {
		fmt.Print(w, "Error: ", view)
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()

	view := gv.Render(w, http.StatusOK, "index", goview.M {
		"title": "Books list",
		"books": books,
	})

	if view != nil {
		fmt.Fprint(w, "Error: ", view)
	}
}

func ShowCreateForm(w http.ResponseWriter, r *http.Request) {
	view := gv.Render(w, http.StatusOK, "create", goview.M {
		"title": "Create new book",
	})

	if view != nil {
		fmt.Fprint(w, "Error: ", view)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Fprint(w, "Error : ", err)
	}

	formData := r.Form

	_, err = models.CreateBook(formData)

	if err != nil {
		fmt.Fprint(w, "Error: ", err)
	}

	http.Redirect(w, r, "/books", 301)
}

func getBook() {

}

func updateBook() {

}

func deleteBook() {

}