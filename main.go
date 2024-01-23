package main

import (
	"BookStoreApi-Go/Controller"
	"BookStoreApi-Go/Middleware"
	"BookStoreApi-Go/Services"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	Services.Init()

	fmt.Printf("Hello BookStore Api\n")
	r := chi.NewRouter()
	r.Use(AddHeaders)

	r.Group(func(r chi.Router) {
		r.Use(Middleware.JwtAuth)
		r.Post("/create", Controller.Create)
		r.Get("/books/{id}", Controller.GetBookById)
		r.Get("/books", Controller.GetAllBooks)
		r.Delete("/books/{id}", Controller.DeleteBookById)
		r.Put("/books/{id}", Controller.UpdateBookById)
	})
	r.Group(func(r chi.Router) {
		r.Use(Middleware.BasicAuth)
		r.Post("/login", Controller.Login)
	})

	http.ListenAndServe("localhost:3000", r)
}

// AddHeaders adds some common header in the response.
func AddHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add {content-type : "application/json"}
		w.Header().Add("content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}
