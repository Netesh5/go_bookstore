package main

import (
	"book_store/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.BookstoreRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe(":9010", r)
	if err != nil {
		log.Fatal(err)
	}
}
