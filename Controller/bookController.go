package Controller

import (
	"BookStoreApi-Go/Model"
	"BookStoreApi-Go/Services"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func Create(res http.ResponseWriter, req *http.Request) {
	var book Model.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	book.UUID = uuid.New()
	Services.IsUuidExist[book.UUID] = true
	Services.SaveBook(book, res)

	fmt.Fprint(res, "Book created successfully\n")
	res.WriteHeader(http.StatusCreated)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	id1 := chi.URLParam(req, "id")
	id, _ := uuid.Parse(id1)
	if Services.IsExist(id) != true {
		http.Error(res, "Book Id does not exist\n", http.StatusNotFound)
		return
	}
	Services.GetBookById(id, res)
}

func DeleteBookById(res http.ResponseWriter, req *http.Request) {
	id1 := chi.URLParam(req, "id")
	id, _ := uuid.Parse(id1)
	if Services.IsExist(id) != true {
		http.Error(res, "Book Id does not exist\n", http.StatusNotFound)
		return
	}
	Services.DeleteBookById(id, res)
	fmt.Fprint(res, "Book deleted successfully\n")
	res.WriteHeader(http.StatusOK)
}

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	Services.GetAllBooks(res)
	res.WriteHeader(http.StatusOK)
}

func UpdateBookById(res http.ResponseWriter, req *http.Request) {
	id1 := chi.URLParam(req, "id")
	id, _ := uuid.Parse(id1)
	if Services.IsExist(id) != true {
		http.Error(res, "Book Id does not exist\n", http.StatusNotFound)
		return
	}
	var book Model.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	Services.UpdateBookById(id, book, res)
	fmt.Fprint(res, "Book updated successfully\n")
	res.WriteHeader(http.StatusOK)
}
