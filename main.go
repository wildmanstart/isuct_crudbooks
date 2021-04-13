package main

import (
	"fmt"
	"main/pkg/routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BookRoutes(r)

	http.Handle("/", r)
	fmt.Print("Listen to http://localhost:9191")
	log.Fatal(http.ListenAndServe("localhost:9191", r))
}
