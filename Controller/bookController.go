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

	jsonBook := Services.SaveBook(book)
	fmt.Fprint(res, "Book created successfully\n")
	res.Write(jsonBook)
	res.WriteHeader(http.StatusCreated)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	id1 := chi.URLParam(req, "id")
	id, _ := uuid.Parse(id1)
	if Services.IsExist(id) != true {
		http.Error(res, "Book Id does not exist\n", http.StatusNotFound)
		return
	}
	jsonBook := Services.GetBookById(id)
	res.Write(jsonBook)
	res.WriteHeader(http.StatusOK)
}

func DeleteBookById(res http.ResponseWriter, req *http.Request) {
	id1 := chi.URLParam(req, "id")
	id, _ := uuid.Parse(id1)
	if Services.IsExist(id) != true {
		http.Error(res, "Book Id does not exist\n", http.StatusNotFound)
		return
	}
	jsonBook := Services.DeleteBookById(id)
	fmt.Fprint(res, "Book deleted successfully\n")
	res.Write(jsonBook)
	res.WriteHeader(http.StatusOK)
}

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	jsonBook := Services.GetAllBooks()
	res.Write(jsonBook)
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
	jsonBook := Services.UpdateBookById(id, book)
	fmt.Fprint(res, "Book updated successfully\n")
	res.Write(jsonBook)
	res.WriteHeader(http.StatusOK)
}
