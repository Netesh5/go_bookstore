package controllers

import (
	"book_store/pkg/models"
	"book_store/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var newBooks models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	setHeader(w)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.NotFound(w, r)
	}
	book, _ := models.GetBook(ID)
	res, _ := json.Marshal(book)
	setHeader(w)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := models.CreateBook(CreateBook)

	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Book couldn't be created", http.StatusBadRequest)
	}
	setHeader(w)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	book := models.DeleteBook(id)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Book couldn't be deleted", http.StatusExpectationFailed)
	}
	setHeader(w)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	book, db := models.GetBook(id)
	if UpdateBook.Name != "" {
		book.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		book.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		book.Publication = UpdateBook.Publication
	}

	db.Save(&book)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Book couldn't be update", http.StatusExpectationFailed)
	}
	setHeader(w)
	w.Write(res)

}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
