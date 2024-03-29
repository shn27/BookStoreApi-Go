package Routes

import (
	"BookStoreApi-Go/Controller"
	"BookStoreApi-Go/Middleware"
	"fmt"
	chiprometheus "github.com/edmarfelipe/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

func Start(port int) {
	fmt.Printf("Hello BookStore Api\n")
	r := chi.NewRouter()
	r.Use(Middleware.AddHeaders)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(chiprometheus.NewMiddleware("service_name"))
	r.Handle("/metrics", promhttp.Handler())

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

	http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), r)
}
