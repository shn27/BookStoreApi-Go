package main

import (
	"BookStoreApi-Go/Controller"
	"BookStoreApi-Go/Services"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	Services.Init()

	fmt.Printf("Hello BookStore Api\n")
	r := chi.NewRouter()
	r.Post("/create", Controller.Create)
	r.Get("/books/{id}", Controller.GetBookById)
	r.Get("/books", Controller.GetAllBooks)
	r.Delete("/books/{id}", Controller.DeleteBookById)
	r.Put("/books/{id}", Controller.UpdateBookById)

	http.ListenAndServe(":8000", r)
}
