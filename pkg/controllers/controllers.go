package controllers

import (
	"fmt"
	"github.com/foolin/goview"
	"main/pkg/models"
	"net/http"
	"main/pkg/utils"
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

func GetBook(w http.ResponseWriter, r *http.Request) {
	book, err := models.GetBook(w, r)

	if err != nil {
		utils.ErrorHandle(w, r, http.StatusNotFound)
		return
	}

	view := gv.Render(w, http.StatusOK, "show", goview.M {
		"title": "Книга " + string(book.Id),
		"book": book,
	})

	if view != nil {
		fmt.Fprint(w, "Error render show page")
	}
}

func ShowUpdateForm(w http.ResponseWriter, r *http.Request) {

	book, err := models.GetBook(w, r)

	if err != nil {
		utils.ErrorHandle(w, r, http.StatusNotFound)
		return
	}

	view := gv.Render(w, http.StatusOK, "update", goview.M {
		"title": "Форма обновления книги \"" + book.Name + "\"",
		"book": book,
	})

	if view != nil {
		fmt.Fprint(w, "Error render show page", view)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Fprint(w, "Error: ", err)
	}

	formData := r.Form

	_, err = models.UpdateBook(formData, w, r)

	if err != nil {
		fmt.Fprint(w, "Error: ", err)
	}

	http.Redirect(w, r, "/books", 301)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	status := models.DeleteBook(w, r)

	if status == true {
		http.Redirect(w, r, "/books", 301)
	} else {
		http.Redirect(w, r, "/books", 301)
	}
}