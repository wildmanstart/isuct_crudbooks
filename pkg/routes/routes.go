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
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET") // Получить книгу (по ее id)
	router.HandleFunc("/book/update/{id}", controllers.ShowUpdateForm).Methods("GET") // Показать форму обновления книги
	router.HandleFunc("/book/update/{id}", controllers.UpdateBook).Methods("POST") // Обновить книгу (по ее id)
	router.HandleFunc("/book/delete/{id}", controllers.DeleteBook).Methods("POST") // Удалить книгу (по ее id)
}