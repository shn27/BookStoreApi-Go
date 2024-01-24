package Controller

import (
	"BookStoreApi-Go/Model"
	"BookStoreApi-Go/Services"
	"encoding/json"
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

	res.WriteHeader(http.StatusCreated)
	Services.SaveBook(book, res)
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
}

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	Services.GetAllBooks(res)
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
}
