package routes

import (
	"main/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookRoutes = func (router *mux.Router) {
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET") // Получить список всех книг
	router.HandleFunc("/book", controllers.ShowCreateForm).Methods("GET")  // Показать страницу с формой создания новой книге
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST") // Создать новую книгу
	//router.HandleFunc("/book/{id}", controllers.getBook).Methods("GET") // Получить книгу (по ее id)
	//router.HandleFunc("/book/{id}", controllers.updateBook).Methods("PUT") // Обновить книгу (по ее id)
	//router.HandleFunc("/book/{id}", controllers.deleteBook).Methods("DELETE") // Удалить книгу (по ее id)
}