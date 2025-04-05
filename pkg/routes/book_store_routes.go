package routes

import (
	"book_store/pkg/controllers"

	"github.com/gorilla/mux"
)

var BookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
