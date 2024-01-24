package Routes

import (
	"BookStoreApi-Go/Controller"
	"BookStoreApi-Go/Middleware"
	"BookStoreApi-Go/Services"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Start() {
	Services.Init()

	fmt.Printf("Hello BookStore Api\n")
	r := chi.NewRouter()
	r.Use(Middleware.AddHeaders)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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
