package main

import (
	"book_store/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.BookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Serving Running on port : 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
