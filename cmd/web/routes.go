package main

import (
	"github.com/ekateryna-tln/wallester_task/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(LoadSession)
	mux.Use(NoSurf)
	mux.Use(URLHandler)

	mux.Get("/{locale}", handlers.Repo.ShowHomePage)
	mux.Get("/{locale}/customers", handlers.Repo.ShowAllCustomers)
	mux.Post("/{locale}/customers/search", handlers.Repo.SearchCustomers)

	mux.Get("/{locale}/customer/{id}/view", handlers.Repo.ShowCustomer)
	mux.Get("/{locale}/customer/{id}", handlers.Repo.ShowCustomerForm)
	mux.Post("/{locale}/customer/{id}", handlers.Repo.EditCustomer)

	mux.Get("/{locale}/customer", handlers.Repo.ShowCustomerForm)
	mux.Post("/{locale}/customer", handlers.Repo.AddCustomer)

	mux.Get("/*", handlers.Repo.PageNotFound)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
