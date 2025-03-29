package routes

import "github.com/gorilla/mux"

var BookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.getBooks).Methods("GET")
	router.HandleFunc("/books/", controllers.createBook).Methods("POST")
	router.HandleFunc("/books/{bookId}", controllers.getBook).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.deleteBook).Methods("DELETE")
}
